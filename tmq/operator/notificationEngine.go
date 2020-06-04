package operator

import (
	"../buffers"
	"../marshaller"
	"../protocol"
	"fmt"
	"log"
)
import "../infrastructure"

type NotificationEngine struct {
	ServerHost string
	ServerPort int
	srh infrastructure.ServerRequestHandler
}

func (ne *NotificationEngine) Initialize() {
	ne.srh = infrastructure.ServerRequestHandler{
				ServerHost:ne.ServerHost,
				ServerPort:ne.ServerPort}

	ne.srh.Initialize()
	go ne.getNotifications()
}

func (ne *NotificationEngine) getNotifications() {
	for {
		topicName := <- buffers.ToNotifyTopicNames
		go ne.keepSending(topicName)
	}
}

func (ne *NotificationEngine) keepSending(topicName string) {
	for {
		messageToSend := <- buffers.Topics[topicName]
		topicSubscribers := buffers.Subscribers[topicName]
		ne.sendToAll(messageToSend, topicSubscribers)
	}
}

func (ne *NotificationEngine) sendToAll(message float64, subscribers []string) {
	m := marshaller.Marshaller{}
	for _, subscriber := range subscribers {
		log.Println(fmt.Sprintf("Sending to conn %s msg %f", subscriber, message))

		packet := protocol.Packet{"message", []interface{}{message}}
		serializedPacket := m.Marshall(packet)

		ne.srh.Send(serializedPacket, subscriber)
	}
}