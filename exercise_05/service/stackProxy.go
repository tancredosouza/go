package service

import "../distribution"

type StackProxy struct {
	HostIp         string
	HostPort       int
	RemoteObjectId int
	TypeName string
}

func (s StackProxy) RemoveElement() string {
	inv := distribution.Requester{}

	res := inv.Invoke(s.HostIp, s.HostPort, s.RemoteObjectId, "pop", []interface{}{})

	return res[0].(string)
}

func (s StackProxy) InsertElement(v int) string {
	inv := distribution.Requester{}

	res := inv.Invoke(s.HostIp, s.HostPort, s.RemoteObjectId, "push", []interface{}{v})

	return res[0].(string)
}

func (s StackProxy) GetFirstElement() string {
	inv := distribution.Requester{}

	res := inv.Invoke(s.HostIp, s.HostPort, s.RemoteObjectId, "top", []interface{}{})

	return res[0].(string)
}

func (s StackProxy) GetSize() string {
	inv := distribution.Requester{}

	res := inv.Invoke(s.HostIp, s.HostPort, s.RemoteObjectId, "size", []interface{}{})

	return res[0].(string)
}