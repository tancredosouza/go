package service

import "../distribution"

type QueueProxy struct {
	ProxyName string
	HostIp         string
	HostPort       int
	RemoteObjectId int
	TypeName string
}

func (q QueueProxy) RemoveElement() string {
	inv := distribution.Requester{}

	res := inv.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "pop", []interface{}{})

	return res
}

func (q QueueProxy) InsertElement(v int) string {
	inv := distribution.Requester{}

	res := inv.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "push", []interface{}{v})

	return res
}

func (q QueueProxy) GetFirstElement() string {
	inv := distribution.Requester{}

	res := inv.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "front", []interface{}{})

	return res
}

func (q QueueProxy) GetSize() string {
	inv := distribution.Requester{}

	res := inv.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "size", []interface{}{})

	return res
}