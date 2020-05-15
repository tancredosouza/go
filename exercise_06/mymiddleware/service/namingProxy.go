package service

import (
	"fmt"
	"github.com/my/repo/mymiddleware/distribution"
	"log"
)

type NamingServiceProxy struct {
	NamingServiceIp   string
	NamingServicePort int

	namingProxyRequester distribution.Requester
}

func (n *NamingServiceProxy) Initialize() {
	n.namingProxyRequester = distribution.Requester{}
	n.namingProxyRequester.Initialize(n.NamingServiceIp, n.NamingServicePort)
}

func (n NamingServiceProxy) Register(proxyName string, proxy Proxy) string {
	res, err := n.namingProxyRequester.Invoke(
		n.NamingServiceIp,
		n.NamingServicePort,
		0,
		"register",
		[]interface{}{proxyName, proxy})

	if (err != nil) {
		log.Fatal(fmt.Sprintf("An error occurred during registration: %s", res))
	}

	n.namingProxyRequester.Conn.Close()

	return res[0].(string)
}

func (n NamingServiceProxy) Lookup(proxyName string) *QueueProxy {
	res, err := n.namingProxyRequester.Invoke(
		n.NamingServiceIp,
		n.NamingServicePort,
		0,
		"lookup",
		[]interface{}{proxyName})
	if (err != nil) {
		log.Fatal("Lookup error. ", res)
	}

	mappedProxy := res[0].(map[string]interface{})
	n.namingProxyRequester.Conn.Close()

	return &(QueueProxy{
		HostIp: mappedProxy["HostIp"].(string),
		HostPort: int(mappedProxy["HostPort"].(float64)),
		RemoteObjectId:int(mappedProxy["RemoteObjectId"].(float64)),
		TypeName: mappedProxy["TypeName"].(string),
		QueueNumber: int(mappedProxy["QueueNumber"].(float64))})
}