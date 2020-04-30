package service

import (
	"../distribution"
	"fmt"
	"log"
)

type NamingServiceProxy struct {
	NamingServiceIp   string
	NamingServicePort int
}

func (n NamingServiceProxy) Register(proxyName string, proxy Proxy) string {
	res, err := distribution.Requester{}.Invoke(
		n.NamingServiceIp,
		n.NamingServicePort,
		0,
		"register",
		[]interface{}{proxyName, proxy})

	if (err != nil) {
		log.Fatal(fmt.Sprintf("An error occurred during registration: %s", res))
	}

	return res[0].(string)
}

func (n NamingServiceProxy) Lookup(proxyName string) Proxy {
	res, err := distribution.Requester{}.Invoke(
		n.NamingServiceIp,
		n.NamingServicePort,
		0,
		"lookup",
		[]interface{}{proxyName})

	if (err != nil) {
		log.Fatal("Lookup error. ", res)
	}

	mappedProxy := res[0].(map[string]interface{})

	if (mappedProxy["TypeName"] == "queue") {
		return QueueProxy{
			mappedProxy["HostIp"].(string),
			int(mappedProxy["HostPort"].(float64)),
			int(mappedProxy["RemoteObjectId"].(float64)),
			mappedProxy["TypeName"].(string)}
	} else {
		return StackProxy{
			mappedProxy["HostIp"].(string),
			int(mappedProxy["HostPort"].(float64)),
			int(mappedProxy["RemoteObjectId"].(float64)),
			mappedProxy["TypeName"].(string)}
	}
}