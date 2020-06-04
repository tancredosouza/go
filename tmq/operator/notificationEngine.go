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
	m := marshaller.Marshaller{}
	for {
		messageToSend := <- buffers.Topics[topicName]
		subscribersOf := buffers.Subscribers[topicName]
		for _, subscriber := range subscribersOf {
			log.Println(fmt.Sprintf("Sending to conn %s msg %f", subscriber, messageToSend))
			ne.srh.Send(m.Marshall(protocol.Packet{"sub_answer", []interface{}{messageToSend}}), subscriber)
		}
	}
}