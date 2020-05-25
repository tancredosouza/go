package service

import (
	"github.com/my/repo/mymiddleware/distribution"
)

type NamingServiceProxy struct {
	HostIp    string
	HostPort  int
	requestor distribution.Requestor
}

func (n *NamingServiceProxy) Initialize() {
	n.requestor = distribution.Requestor{}
	n.requestor.Initialize(n.HostIp, n.HostPort)
}

func (n *NamingServiceProxy) Register(proxyName string, proxy Proxy) string {
	n.requestor.Invoke(
		n.HostIp,
		n.HostPort,
		0,
		"register",
		[]interface{}{proxyName, proxy})

	res := n.requestor.Receive()

	// n.requestor.Crh.CloseConnection()

	return res[0].(string)
}

func (n *NamingServiceProxy) Lookup(proxyName string) *QueueProxy {
	n.requestor.Invoke(
		n.HostIp,
		n.HostPort,
		0,
		"lookup",
		[]interface{}{proxyName})

	res := n.requestor.Receive()

	mappedProxy := res[0].(map[string]interface{})

	// n.requestor.Crh.CloseConnection()

	return &(QueueProxy{
		HostIp: mappedProxy["HostIp"].(string),
		HostPort: int(mappedProxy["HostPort"].(float64)),
		RemoteObjectId:int(mappedProxy["RemoteObjectId"].(float64)),
		TypeName: mappedProxy["TypeName"].(string),
		QueueNumber: int(mappedProxy["QueueNumber"].(float64))})
}