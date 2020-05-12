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
	QueueNumber int
}

func (q QueueProxy) RemoveElement() string {
	requester := distribution.Requester{}

	res, err := requester.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "pop", []interface{}{q.QueueNumber})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}

func (q QueueProxy) InsertElement(v int) string {
	requester := distribution.Requester{}

	res, err := requester.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "push", []interface{}{q.QueueNumber, v})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}

func (q QueueProxy) GetFirstElement() string {
	requester := distribution.Requester{}

	res, err := requester.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "front", []interface{}{q.QueueNumber})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}

func (q QueueProxy) GetSize() string {
	requester := distribution.Requester{}

	res, err := requester.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "size", []interface{}{q.QueueNumber})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}