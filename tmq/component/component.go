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

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (c *Component) Dial(serverHost string, serverPort int) {
	c.id = c.FetchComponentId()

	c.crh = infrastructure.ClientRequestHandler{
		serverHost,
		serverPort}

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
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (c *Component) register() {
	c.marshallPacketAndSend(
		protocol.Packet{"register", []interface{}{c.id} })
}

func (c *Component) Subscribe(topicName string) {
	c.marshallPacketAndSend(
		protocol.Packet{"subscribe", []interface{}{c.id, topicName} })
}

func (c *Component) Publish(topicName string, message string) {
	c.marshallPacketAndSend(
		protocol.Packet{"publish", []interface{}{c.id, topicName, message} })
}

func (c *Component) Unsubscribe(topicName string) {
	c.marshallPacketAndSend(
		protocol.Packet{"unsubscribe", []interface{}{c.id, topicName} })
}

func (c *Component) marshallPacketAndSend(packetToSend protocol.Packet) {
	m := marshaller.Marshaller{}
	bytesToSend := m.Marshall(packetToSend)
	c.crh.SendAndReceive(bytesToSend)
}