package component

import (
	"../infrastructure"
	"../marshaller"
	"../protocol"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

type Component struct{
	Key int
	id  string
	crh infrastructure.ClientRequestHandler
}

func (c *Component) TmqConnect(serverHost string, serverPort int) {
	c.id = c.FetchComponentId()

	c.crh = infrastructure.ClientRequestHandler{
		ServerHost: serverHost,
		ServerPort: serverPort}
	c.crh.Initialize()

	c.register()
}

func (c *Component) FetchComponentId() string {
	filepath := fmt.Sprintf("component/database/myid_%d.txt", c.Key)
	content, err := ioutil.ReadFile(filepath)
	if err == nil {
		return string(content)
	}

	id := RandStringBytes(4)

	f, err := os.Create(filepath)
	if (err != nil) {
		log.Fatal("Error creating file", err)
	}
	_, err = f.WriteString(id)
	if (err != nil) {
		log.Fatal("Error writing id", err)
	}

	return id
}

func RandStringBytes(n int) string {
	const LETTER_BYTES = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = LETTER_BYTES[rand.Intn(len(LETTER_BYTES))]
	}
	return string(b)
}

func (c *Component) register() {
	c.serializeAndSend(
		protocol.Packet{"register", []interface{}{c.id} })
}

func (c *Component) CreateTopic(topicName string) {
	c.serializeAndSend(
		protocol.Packet{"create", []interface{}{c.id, topicName} })
}

func (c *Component) Subscribe(topicName string) {
	c.serializeAndSend(
		protocol.Packet{"subscribe", []interface{}{c.id, topicName} })
}

func (c *Component) Publish(topicName string, message string) {
	c.serializeAndSend(
		protocol.Packet{"publish", []interface{}{c.id, topicName, message} })
}

func (c *Component) Unsubscribe(topicName string) {
	c.serializeAndSend(
		protocol.Packet{"unsubscribe", []interface{}{c.id, topicName} })
}

func (c *Component) serializeAndSend(packetToSend protocol.Packet) {
	m := marshaller.Marshaller{}
	bytesToSend := m.Marshall(packetToSend)

	c.crh.Send(bytesToSend)
}

func (c *Component) receiveAndDeserialize() protocol.Packet	{
	m := marshaller.Marshaller{}
	receivedBytes := c.crh.Receive()

	return m.Unmarshall(receivedBytes)
}