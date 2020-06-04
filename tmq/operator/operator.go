package operator

import (
	"../buffers"
	"../marshaller"
	"../protocol"
	"fmt"
	"log"
)

type Operator struct{}

func Initialize() {
	go KeepListeningToIncomingMessages()
}

func KeepListeningToIncomingMessages() {
	for {
		msg := <- buffers.IncomingMessages
		operate(msg)
	}
}

func operate(msg []byte) {
	m := marshaller.Marshaller{}
	packet := m.Unmarshall(msg)

	switch packet.Operation {
	case "create":
		createTopic(packet.Params[1].(string))
	case "publish":
		publish(packet)
	case "subscribe":
		subscribe(packet)
	}
}

func createTopic(name string) {
	if _, found := buffers.Topics[name]; !found {
		buffers.Topics[name] = make(chan float64, 100)
		buffers.ToNotifyTopicNames <- name
		log.Println(fmt.Sprintf("Created topic with name %s!", name))
	}
}

func publish(p protocol.Packet) {
	topicName := p.Params[1].(string)
	message := p.Params[2].(float64)

	buffers.Topics[topicName] <- message
}

func subscribe(p protocol.Packet) {
	connId := p.Params[0].(string)
	topicName := p.Params[1].(string)

	if (!isAlreadySubscribed(connId, topicName)) {
		buffers.Subscribers[topicName] = append(buffers.Subscribers[topicName], connId)
	}
}

func isAlreadySubscribed(connId string, topicName string) bool {
	for _,elem := range buffers.Subscribers[topicName] {
		if elem == connId {
			return true
		}
	}
	return false
}