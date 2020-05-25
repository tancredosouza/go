package service

import (
	"github.com/my/repo/mymiddleware/distribution"
	"github.com/my/repo/mymiddleware/result_callback"
)

type QueueProxy struct {
	HostIp         string
	HostPort       int
	RemoteObjectId int
	TypeName string
	QueueNumber int

	requestor distribution.Requestor
}

func (q *QueueProxy) Initialize(resultCallback *result_callback.ResultCallback) {
	q.requestor = distribution.Requestor{}
	q.requestor.Initialize(q.HostIp, q.HostPort)

	// This function is defined exclusively for the queueProxy, as the namingProxy is not async.
	// Thus, moving this function into requestor.Initialize() would cause the Server to loop
	// infinitely when waiting for naming proxy to return (it never adds anything to its queue)
	go q.requestor.WaitForResponseAsync(resultCallback)
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