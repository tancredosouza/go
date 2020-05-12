package namingService

import (
	"errors"
	"fmt"
	"github.com/my/repo/mymiddleware/constants"
	"github.com/my/repo/mymiddleware/infrastructure"
	"github.com/my/repo/mymiddleware/marshaller"
	"github.com/my/repo/mymiddleware/protocol"
	"github.com/my/repo/mymiddleware/service"
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


	srh.StartListening()
	for {
		srh.AcceptNewConnection()
		go func() {
			receivedData, _ := srh.Receive()

			processedData := demuxAndProcess(receivedData)

			srh.Send(processedData)
			srh.CloseConnection()
		}()
	}

	srh.StopListening()
}

func demuxAndProcess(data []byte) []byte {
	m := marshaller.Marshaller{}
	p := m.Unmarshall(data)

	proxyName := p.Body.RequestBody.Data[0].(string)
	operation := p.Body.RequestHeader.Operation

	var responseBody protocol.ResponseBody
	var statusCode int
	switch operation {
	case "lookup":
		responseBody, statusCode = lookupAndPack(proxyName)
		break;
	case "register":
		assembledProxy := assembleProxyFromPacket(p)
		err := namingService.registerProxy(assembledProxy, proxyName)
		if err != nil {
			responseBody = protocol.ResponseBody{Data: []interface{}{err.Error()}}
			statusCode = constants.INTERNAL_ERROR
		} else {
			responseBody = protocol.ResponseBody{Data: []interface{}{"Successfully registered!"}}
			statusCode = constants.OK_STATUS
		}
		break;
	default:
		responseBody = protocol.ResponseBody{Data: []interface{}{fmt.Sprintf("Invalid operation %s!", operation)}}
		statusCode = constants.INTERNAL_ERROR
	}

	responseHeader := protocol.ResponseHeader{RequestId: p.Body.RequestHeader.RequestId, Status: statusCode}
	packet := assemblePacket(responseHeader, responseBody)
	serializedPacket := m.Marshall(packet)
	return serializedPacket
}

func lookupAndPack(proxyName string) (protocol.ResponseBody, int) {
	proxy, err := namingService.lookup(proxyName)

	if err != nil {
		return protocol.ResponseBody{Data: []interface{}{}}, constants.NOT_FOUND_STATUS
	} else {
		return protocol.ResponseBody{Data: []interface{}{proxy}}, constants.OK_STATUS
	}
}

func assembleProxyFromPacket(p protocol.Packet) service.Proxy {
	data := p.Body.RequestBody.Data[1].(map[string]interface{})
	hostIp := data["HostIp"].(string)
	hostPort := int(data["HostPort"].(float64))
	proxyType := data["TypeName"].(string)
	if proxyType == constants.QUEUE_TYPE {
		queueNumber := int(data["QueueNumber"].(float64))

		return service.QueueProxy{
				HostIp:         hostIp,
				HostPort:       hostPort,
				TypeName:       proxyType,
				RemoteObjectId: constants.QUEUE_ID,
				QueueNumber: queueNumber,
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

func assemblePacket(responseHeader protocol.ResponseHeader, responseBody protocol.ResponseBody) protocol.Packet {
	body := protocol.Body{ResponseHeader: responseHeader, ResponseBody: responseBody}
	header := protocol.Header{
		Magic: "IF711",
		Version: "1.0",
	}

	packet := protocol.Packet{header,body}
	return packet
}


func (n NamingService) listProxies() {
	// TODO IMPLEMENT FUNCTION
	for name, proxy := range n.data {
		fmt.Sprintf("[%s -> %s]\n", name, proxy)
	}
}
