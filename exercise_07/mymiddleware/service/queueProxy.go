package service

import (
	"github.com/my/repo/mymiddleware/distribution"
)

type QueueProxy struct {
	HostIp         string
	HostPort       int
	RemoteObjectId int
	TypeName string
	QueueNumber int

	requester distribution.Requestor
}

func (q *QueueProxy) Initialize() {
	q.requester = distribution.Requestor{}
	q.requester.Initialize(q.HostIp, q.HostPort)
}
/*
func (q QueueProxy) RemoveElement() string {
	res, err := q.requester.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "pop", []interface{}{q.QueueNumber})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}

func (q QueueProxy) InsertElement(v int) string {
	res, err := q.requester.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "push", []interface{}{q.QueueNumber, v})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}

func (q QueueProxy) GetFirstElement() string {
	res, err := q.requester.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "front", []interface{}{q.QueueNumber})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}

func (q QueueProxy) GetSize() string {
	res, err := q.requester.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "size", []interface{}{q.QueueNumber})
	if (err != nil) {
		log.Fatal(res[0].(string))
	}

	return res[0].(string)
}

 */