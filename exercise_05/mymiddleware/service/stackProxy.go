package service

import (
	"github.com/my/repo/mymiddleware/distribution"
	"log"
)

type StackProxy struct {
	HostIp         string
	HostPort       int
	RemoteObjectId int
	TypeName string
}
var stackRequester *distribution.Requester

func (s StackProxy) RemoveElement() string {
	if (stackRequester == nil) {
		stackRequester = &distribution.Requester{}
	}

	res, err := stackRequester.Invoke(s.HostIp, s.HostPort, s.RemoteObjectId, "pop", []interface{}{})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}

func (s StackProxy) InsertElement(v int) string {
	if (stackRequester == nil) {
		stackRequester = &distribution.Requester{}
	}
	res, err := stackRequester.Invoke(s.HostIp, s.HostPort, s.RemoteObjectId, "push", []interface{}{v})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}

func (s StackProxy) GetFirstElement() string {
	if (stackRequester == nil) {
		stackRequester = &distribution.Requester{}
	}

	res, err := stackRequester.Invoke(s.HostIp, s.HostPort, s.RemoteObjectId, "top", []interface{}{})
	if (err != nil) {
		log.Fatal(res, err)
	}

	return res[0].(string)
}

func (s StackProxy) GetSize() string {
	if (stackRequester == nil) {
		stackRequester = &distribution.Requester{}
	}

	res, err := stackRequester.Invoke(s.HostIp, s.HostPort, s.RemoteObjectId, "size", []interface{}{})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}