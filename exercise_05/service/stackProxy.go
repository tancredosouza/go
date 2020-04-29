package service

import "../distribution"

type StackProxy Proxy

func (s StackProxy) Pop() string {
	inv := distribution.Requester{}

	res := inv.Invoke(s.HostIp, s.Port, s.RemoteObjectId, "pop", []interface{}{})

	return res
}

func (s StackProxy) Push(v int) string {
	inv := distribution.Requester{}

	res := inv.Invoke(s.HostIp, s.Port, s.RemoteObjectId, "push", []interface{}{v})

	return res
}

func (s StackProxy) Top() string {
	inv := distribution.Requester{}

	res := inv.Invoke(s.HostIp, s.Port, s.RemoteObjectId, "top", []interface{}{})

	return res
}

func (s StackProxy) Size() string {
	inv := distribution.Requester{}

	res := inv.Invoke(s.HostIp, s.Port, s.RemoteObjectId, "size", []interface{}{})

	return res
}