package service

import (
	"fmt"
	"github.com/my/repo/mymiddleware/distribution"
	"log"
)

type NamingServiceProxy struct {
	NamingServiceIp   string
	NamingServicePort int
}

var namingProxyRequester *distribution.Requester

func (n NamingServiceProxy) Register(proxyName string, proxy Proxy) string {
	if (namingProxyRequester == nil) {
		//log.Println("Creating requester")
		namingProxyRequester = &distribution.Requester{}
	}

	res, err := namingProxyRequester.Invoke(
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
	if (namingProxyRequester == nil) {
		namingProxyRequester = &distribution.Requester{}
	}
	res, err := namingProxyRequester.Invoke(
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
			HostIp: mappedProxy["HostIp"].(string),
			HostPort: int(mappedProxy["HostPort"].(int64)),
			RemoteObjectId:int(mappedProxy["RemoteObjectId"].(int64)),
			TypeName: mappedProxy["TypeName"].(string),
			QueueNumber: int(mappedProxy["QueueNumber"].(int64))}
	} else {
		return StackProxy{
			HostIp: mappedProxy["HostIp"].(string),
			HostPort: int(mappedProxy["HostPort"].(int64)),
			RemoteObjectId: int(mappedProxy["RemoteObjectId"].(int64)),
			TypeName: mappedProxy["TypeName"].(string)}
	}
}