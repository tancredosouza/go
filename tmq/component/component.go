package component

import (
	"../infrastructure"
	"../marshaller"
	"../protocol"
	"math/rand"
)

type Component struct{
	id string
	crh infrastructure.ClientRequestHandler
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (c *Component) Dial(serverHost string, serverPort int) {
	c.id = RandStringBytes(4)

	c.crh = infrastructure.ClientRequestHandler{
		serverHost,
		serverPort}

	c.Publish(protocol.Packet{"register", []interface{}{c.id} })
}

func (c *Component) Publish(msgToSend protocol.Packet) {
	m := marshaller.Marshaller{}

	bytesToSend := m.Marshall(msgToSend)
	c.crh.SendAndReceive(bytesToSend)
}