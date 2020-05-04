package service

import (
	"github.com/my/repo/mymiddleware/distribution"
	"log"
)

type QueueProxy struct {
	HostIp         string
	HostPort       int
	RemoteObjectId int
	TypeName string
}
var requester *distribution.Requester = nil

func (q QueueProxy) RemoveElement() string {
	if (requester == nil) {
		requester = &distribution.Requester{}
	}

	res, err := requester.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "pop", []interface{}{})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}

func (q QueueProxy) InsertElement(v int) string {
	if (requester == nil) {
		requester = &distribution.Requester{}
	}

	res, err := requester.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "push", []interface{}{v})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}

func (q QueueProxy) GetFirstElement() string {
	if (requester == nil) {
		requester = &distribution.Requester{}
	}

	res, err := requester.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "front", []interface{}{})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}

func (q QueueProxy) GetSize() string {
	if (requester == nil) {
		requester = &distribution.Requester{}
	}

	res, err := requester.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "size", []interface{}{})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}