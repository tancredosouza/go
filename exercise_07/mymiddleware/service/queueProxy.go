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
	go q.requester.ResultCallback()
}

func (q QueueProxy) RemoveElement() {
	q.requester.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "pop", []interface{}{q.QueueNumber})
}

func (q QueueProxy) InsertElement(v int) {
	q.requester.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "push", []interface{}{q.QueueNumber, v})
}

func (q QueueProxy) GetFirstElement() {
	q.requester.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "front", []interface{}{q.QueueNumber})
}

func (q QueueProxy) GetSize() {
	q.requester.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "size", []interface{}{q.QueueNumber})
}