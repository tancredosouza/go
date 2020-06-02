package component

import "../marshaller"
import "../infrastructure"

type Component struct{
	crh infrastructure.ClientRequestHandler
}

func (c *Component) Dial(serverHost string, serverPort int) {
	c.crh = infrastructure.ClientRequestHandler{
		serverHost,
		serverPort}
}

func (c *Component) Publish(msgToSend string) {
	m := marshaller.Marshaller{}

	bytesToSend := m.Marshall(msgToSend)
	c.crh.SendAndReceive(bytesToSend)
}