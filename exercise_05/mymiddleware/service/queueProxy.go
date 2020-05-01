package service

import (
	"../distribution"
	"log"
)

type QueueProxy struct {
	HostIp         string
	HostPort       int
	RemoteObjectId int
	TypeName string
}

func (q QueueProxy) RemoveElement() string {
	inv := distribution.Requester{}

	res, err := inv.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "pop", []interface{}{})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}

func (q QueueProxy) InsertElement(v int) string {
	inv := distribution.Requester{}

	res, err := inv.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "push", []interface{}{v})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}

func (q QueueProxy) GetFirstElement() string {
	inv := distribution.Requester{}

	res, err := inv.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "front", []interface{}{})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}

func (q QueueProxy) GetSize() string {
	inv := distribution.Requester{}

	res, err := inv.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "size", []interface{}{})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}