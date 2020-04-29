package main

import (
	"../service"
	"errors"
	"fmt"
	"../infrastructure"
	"net"
	"../common"
)

var listener net.Listener
var err error
var clientConn net.Conn

type NamingService struct {
	data map[string] service.Proxy
}

func main() {
	n := NamingService{}
	n.data = map[string]service.Proxy{}
	n.StartService("localhost", 3245)
}

func (n NamingService) StartService(ip string, port int) {
	handler := infrastructure.ServerRequestHandler{
		ServerHost: ip,
		ServerPort: port,
	}

	for {
		receivedData := handler.Receive()

		processedData := n.demuxAndProcess(receivedData)

		handler.Send(processedData)
	}
}

func (n NamingService) demuxAndProcess(data []byte) []byte {
	m := common.Marshaller{}
	p := m.Unmarshall(data)

	// choose operation
	i := p.Body.RequestBody.Data[0].(map[string]interface{})

	name := i["ProxyName"].(string)
	op := p.Body.RequestHeader.Operation

	switch op {
	case "lookup":
		s, err := n.lookup(name)

		if err != nil {
			return []byte(fmt.Sprint(err))
		}

		return m.Marshall(s)
		break;
	case "register":
		ip := i["HostIp"].(string)
		t := i["TypeName"].(string)
		port := int(i["HostPort"].(float64))

		var err error
		if t == "queue" {
			err = n.registerProxy(service.QueueProxy{HostIp: ip, HostPort: port, TypeName: t, RemoteObjectId: 2}, name)
		}

		if t == "stack" {
			err = n.registerProxy(service.StackProxy{HostIp: ip, HostPort: port, TypeName: t, RemoteObjectId: 1}, name)
		}

		if err != nil {
			return []byte(fmt.Sprint(err))
		}

		return []byte(fmt.Sprint("Successfuly registered ", name))
	}

	return []byte(fmt.Sprint("Invalid operation %s!", op))
}


func (n NamingService) registerProxy(proxy service.Proxy, proxyName string) error {
	if _, isNameRegistered := n.data[proxyName]; !isNameRegistered {
		n.data[proxyName] = proxy
		return nil
	} else {
		return errors.New(fmt.Sprint("Name %s already registered!", proxyName))
	}
}

func (n NamingService) lookup(proxyName string) (service.Proxy, error) {
	if _, isNameRegistered := n.data[proxyName]; isNameRegistered {
		return n.data[proxyName], nil
	} else {
		return nil, errors.New(fmt.Sprint("Name %s is not registered!", proxyName))
	}
}

func (n NamingService) listProxies() {
	for name, proxy := range n.data {
		fmt.Sprintln("[%s -> %s", name, proxy)
	}
}
