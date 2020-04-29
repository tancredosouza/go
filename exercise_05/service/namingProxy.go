package service

import (
	"../common"
	"../infrastructure"
	"../packetdef"
	"bytes"
	"encoding/json"
	"log"
)

type NamingServiceProxy struct {
	NamingServiceIp   string
	NamingServicePort int
}

func (n NamingServiceProxy) Register(proxy Proxy) string {
	m := common.Marshaller{}

	//assemble packet
	b := packetdef.RequestBody{
		Data: []interface{}{proxy},
	}
	h := packetdef.RequestHeader{Operation:"register"}
	body := packetdef.Body{RequestHeader: h, RequestBody: b}
	header := packetdef.Header{
		Magic: "IF711",
		Version: "1.0",
	}

	packet := packetdef.Packet{header,body}
	data := m.Marshall(packet)

	crh := infrastructure.ClientRequestHandler{ServerHost: n.NamingServiceIp, ServerPort: n.NamingServicePort}
	res := crh.SendAndReceive(data)

	return string(res)
}

// func (n NamingService) lookup(proxyName string) (service.Proxy, error) {
func (n NamingServiceProxy) Lookup(proxyName string) Proxy {
	//assemble packet
	b := packetdef.RequestBody{
		Data: []interface{}{map[string]string{"ProxyName": proxyName}},
	}
	h := packetdef.RequestHeader{Operation:"lookup"}
	body := packetdef.Body{RequestHeader: h, RequestBody: b}
	header := packetdef.Header{
		Magic: "IF711",
		Version: "1.0",
	}

	packet := packetdef.Packet{header,body}
	data, _ := json.Marshal(packet)

	crh := infrastructure.ClientRequestHandler{ServerHost: n.NamingServiceIp, ServerPort: n.NamingServicePort}
	res := crh.SendAndReceive(data)

	var mres map[string]interface{}
	err := json.Unmarshal(bytes.Trim(res, "\x00"), &mres)
	if (err != nil) {
		log.Fatal("Error ocurred ", err)
	}

	if (mres["TypeName"] == "queue") {
		var queueProxy QueueProxy
		json.Unmarshal(bytes.Trim(res, "\x00"), &queueProxy)

		return queueProxy
	} else {
		var stackProxy StackProxy
		json.Unmarshal(bytes.Trim(res, "\x00"), &stackProxy)

		return stackProxy
	}
}