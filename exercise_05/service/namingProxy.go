package service

import (
	"../distribution"
)

type NamingServiceProxy struct {
	NamingServiceIp   string
	NamingServicePort int
}

func (n NamingServiceProxy) Register(proxyName string, proxy Proxy) string {
	res := distribution.Requester{}.Invoke(
		n.NamingServiceIp,
		n.NamingServicePort,
		0,
		"register",
		[]interface{}{proxyName, proxy})

	return res[0].(string)
}

func (n NamingServiceProxy) Lookup(proxyName string) Proxy {
	res := distribution.Requester{}.Invoke(
		n.NamingServiceIp,
		n.NamingServicePort,
		0,
		"lookup",
		[]interface{}{proxyName})[0].(map[string]interface{})

	if (res["TypeName"] == "queue") {
		return QueueProxy{
			res["HostIp"].(string),
			int(res["HostPort"].(float64)),
			int(res["RemoteObjectId"].(float64)),
			res["TypeName"].(string)}
	} else {
		return StackProxy{
			res["HostIp"].(string),
			int(res["HostPort"].(float64)),
			int(res["RemoteObjectId"].(float64)),
			res["TypeName"].(string)}
	}
}