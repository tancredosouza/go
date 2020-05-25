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

	requestor distribution.Requestor
}

func (q *QueueProxy) Initialize() {
	q.requestor = distribution.Requestor{}
	q.requestor.Initialize(q.HostIp, q.HostPort)
	go q.requestor.WaitForResponseAsync()
}

func (q QueueProxy) RemoveElement() {
	q.requestor.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "pop", []interface{}{q.QueueNumber})
}

func (q QueueProxy) InsertElement(v int) {
	q.requestor.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "push", []interface{}{q.QueueNumber, v})
}

func (q QueueProxy) GetFirstElement() {
	q.requestor.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "front", []interface{}{q.QueueNumber})
}

func (q QueueProxy) GetSize() {
	q.requestor.Invoke(q.HostIp, q.HostPort, q.RemoteObjectId, "size", []interface{}{q.QueueNumber})
}