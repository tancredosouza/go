package service

import (
	"../distribution"
	"log"
)

type StackProxy struct {
	HostIp         string
	HostPort       int
	RemoteObjectId int
	TypeName string
}

func (s StackProxy) RemoveElement() string {
	inv := distribution.Requester{}

	res, err := inv.Invoke(s.HostIp, s.HostPort, s.RemoteObjectId, "pop", []interface{}{})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}

func (s StackProxy) InsertElement(v int) string {
	inv := distribution.Requester{}

	res, err := inv.Invoke(s.HostIp, s.HostPort, s.RemoteObjectId, "push", []interface{}{v})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}

func (s StackProxy) GetFirstElement() string {
	inv := distribution.Requester{}

	res, err := inv.Invoke(s.HostIp, s.HostPort, s.RemoteObjectId, "top", []interface{}{})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}

func (s StackProxy) GetSize() string {
	inv := distribution.Requester{}

	res, err := inv.Invoke(s.HostIp, s.HostPort, s.RemoteObjectId, "size", []interface{}{})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}