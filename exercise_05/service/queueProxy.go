package service

import "../distribution"

type QueueProxy Proxy

func (q QueueProxy) Pop() string {
	inv := distribution.Requester{}

	res := inv.Invoke(q.HostIp, q.Port, q.RemoteObjectId, "pop", []interface{}{})

	return res
}

func (q QueueProxy) Push(v int) string {
	inv := distribution.Requester{}

	res := inv.Invoke(q.HostIp, q.Port, q.RemoteObjectId, "push", []interface{}{v})

	return res
}

func (q QueueProxy) Front() string {
	inv := distribution.Requester{}

	res := inv.Invoke(q.HostIp, q.Port, q.RemoteObjectId, "front", []interface{}{})

	return res
}

func (q QueueProxy) Size() string {
	inv := distribution.Requester{}

	res := inv.Invoke(q.HostIp, q.Port, q.RemoteObjectId, "size", []interface{}{})

	return res
}