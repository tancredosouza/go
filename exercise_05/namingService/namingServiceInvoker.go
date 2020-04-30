package namingService

import (
	"../constants"
	"../infrastructure"
	"../marshaller"
	"../packetdef"
	"../service"
	"errors"
	"fmt"
	"log"
)

type NamingService struct {
	data map[string]service.Proxy
}

var namingService NamingService

type Invoker struct{
	HostIp string
	HostPort int
}

func (i Invoker) Invoke() {
	namingService = NamingService{map[string]service.Proxy{}}

	srh := infrastructure.ServerRequestHandler{
		ServerHost: i.HostIp,
		ServerPort: i.HostPort,
	}

	for {
		receivedData := srh.Receive()

		processedData := i.demuxAndProcess(receivedData)

		srh.Send(processedData)
	}
}

func (i Invoker) demuxAndProcess(data []byte) []byte {
	m := marshaller.Marshaller{}
	p := m.Unmarshall(data)

	proxyName := p.Body.RequestBody.Data[0].(string)
	operation := p.Body.RequestHeader.Operation

	var responseBody packetdef.ResponseBody
	switch operation {
	case "lookup":
		responseBody = lookupAndPack(proxyName)
		break;
	case "register":
		assembledProxy := assembleProxyFromPacket(p)
		err := namingService.registerProxy(assembledProxy, proxyName)
		if err != nil {
			responseBody = packetdef.ResponseBody{Data: []interface{}{err}}
		} else {
			responseBody = packetdef.ResponseBody{Data: []interface{}{"Successfully registered!"}}
		}
		break;
	default:
		responseBody = packetdef.ResponseBody{Data: []interface{}{fmt.Sprintf("Invalid operation %s!", operation)}}
	}

	responseHeader := packetdef.ResponseHeader{RequestId: p.Body.RequestHeader.RequestId}
	packet := assemblePacket(responseHeader, responseBody)
	serializedPacket := m.Marshall(packet)
	return serializedPacket
}

func lookupAndPack(proxyName string) packetdef.ResponseBody {
	proxy, err := namingService.lookup(proxyName)

	var responseBody packetdef.ResponseBody
	if err != nil {
		responseBody = packetdef.ResponseBody{Data: []interface{}{proxy}}
	}

	responseBody = packetdef.ResponseBody{Data: []interface{}{proxy}}
	return responseBody
}

func assembleProxyFromPacket(p packetdef.Packet) service.Proxy {
	data := p.Body.RequestBody.Data[1].(map[string]interface{})
	hostIp := data["HostIp"].(string)
	hostPort := int(data["HostPort"].(float64))
	proxyType := data["TypeName"].(string)

	if proxyType == constants.QUEUE_TYPE {
			return service.QueueProxy{
				HostIp:         hostIp,
				HostPort:       hostPort,
				TypeName:       proxyType,
				RemoteObjectId: constants.QUEUE_ID,
			}
	}

	if proxyType == constants.STACK_TYPE {
		return service.StackProxy{
				HostIp:         hostIp,
				HostPort:       hostPort,
				TypeName:       proxyType,
				RemoteObjectId: constants.STACK_ID,
		}
	}

	log.Panic("Invalid proxyType ", proxyType)
	return nil
}

func (n NamingService) lookup(proxyName string) (service.Proxy, error) {
	if _, isNameRegistered := n.data[proxyName]; isNameRegistered {
		return n.data[proxyName], nil
	} else {
		return nil, errors.New(fmt.Sprintf("Name %s is not registered!", proxyName))
	}
}

func (n NamingService) registerProxy(proxy service.Proxy, proxyName string) error {
	if _, isNameRegistered := n.data[proxyName]; !isNameRegistered {
		n.data[proxyName] = proxy
		return nil
	} else {
		return errors.New(fmt.Sprintf("Name %s already registered!", proxyName))
	}
}

func assemblePacket(responseHeader packetdef.ResponseHeader, responseBody packetdef.ResponseBody) packetdef.Packet {
	body := packetdef.Body{ResponseHeader: responseHeader, ResponseBody: responseBody}
	header := packetdef.Header{
		Magic: "IF711",
		Version: "1.0",
	}

	packet := packetdef.Packet{header,body}
	return packet
}


func (n NamingService) listProxies() {
	// TODO IMPLEMENT FUNCTION
	for name, proxy := range n.data {
		fmt.Sprintf("[%s -> %s]\n", name, proxy)
	}
}
