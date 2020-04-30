package namingService

import (
	"../constants"
	"../infrastructure"
	"../marshaller"
	"../packetdef"
	"../service"
	"errors"
	"fmt"
)

type NamingService struct {
	data map[string]service.Proxy
}

var namingService NamingService

type NamingServiceInvoker struct{
	HostIp string
	HostPort int
}

func (i NamingServiceInvoker) Invoke() {
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

func (NamingServiceInvoker) demuxAndProcess(data []byte) []byte {
	m := marshaller.Marshaller{}
	p := m.Unmarshall(data)

	name := p.Body.RequestBody.Data[0].(string)
	op := p.Body.RequestHeader.Operation

	switch op {
	case "lookup":
		proxy, err := namingService.lookup(name)

		if err != nil {
			return []byte(fmt.Sprint(err))
		}

		b := packetdef.ResponseBody{
			Data: []interface{}{proxy},
		}
		h := packetdef.ResponseHeader{RequestId: p.Body.RequestHeader.RequestId}
		body := packetdef.Body{ResponseHeader: h, ResponseBody: b}
		header := packetdef.Header{
			Magic: "IF711",
			Version: "1.0",
		}

		packet := packetdef.Packet{header,body}

		return m.Marshall(packet)
	case "register":
		i := p.Body.RequestBody.Data[1].(map[string]interface{})
		ip := i["HostIp"].(string)
		t := i["TypeName"].(string)
		port := int(i["HostPort"].(float64))

		var err error
		var res string
		if t == constants.QUEUE_TYPE {
			err = namingService.registerProxy(
				service.QueueProxy{
					HostIp: ip,
					HostPort: port,
					TypeName: t,
					RemoteObjectId: constants.QUEUE_ID}, name)
		}

		if t == constants.STACK_TYPE {
			err = namingService.registerProxy(
				service.StackProxy{
					HostIp: ip,
					HostPort: port,
					TypeName: t,
					RemoteObjectId: constants.STACK_ID}, name)
		}

		if err != nil {
			res = fmt.Sprint(err)
		} else {
			res = "Successfully registered!"
		}

		b := packetdef.ResponseBody{
			Data: []interface{}{res},
		}
		h := packetdef.ResponseHeader{RequestId: p.Body.RequestHeader.RequestId}
		body := packetdef.Body{ResponseHeader: h, ResponseBody: b}
		header := packetdef.Header{
			Magic: "IF711",
			Version: "1.0",
		}

		packet := packetdef.Packet{header,body}

		return m.Marshall(packet)
	}

	return []byte(fmt.Sprintf("Invalid operation %s!", op))
}


func (n NamingService) registerProxy(proxy service.Proxy, proxyName string) error {
	if _, isNameRegistered := n.data[proxyName]; !isNameRegistered {
		n.data[proxyName] = proxy
		return nil
	} else {
		return errors.New(fmt.Sprintf("Name %s already registered!", proxyName))
	}
}

func (n NamingService) lookup(proxyName string) (service.Proxy, error) {
	if _, isNameRegistered := n.data[proxyName]; isNameRegistered {
		return n.data[proxyName], nil
	} else {
		return nil, errors.New(fmt.Sprintf("Name %s is not registered!", proxyName))
	}
}

func (n NamingService) listProxies() {
	for name, proxy := range n.data {
		fmt.Sprintf("[%s -> %s]\n", name, proxy)
	}
}
