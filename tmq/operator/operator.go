package operator

import (
	"../buffers"
	"../locks"
	"../marshaller"
	"../persist"
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

		locks.TopicsLock.Lock()
		buffers.Topics[name] = make(chan float64, 100)
		locks.TopicsLock.Unlock()
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

	locks.SubscribersLock.Lock()
	if (!isAlreadySubscribed(connId, topicName)) {
		buffers.Subscribers[topicName] = append(buffers.Subscribers[topicName], connId)

		err := persist.Save("./database/subscribers", buffers.Subscribers)
		locks.SubscribersLock.Unlock()
		if (err != nil) {
			log.Fatal("Error persisting database ", err)
		}
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