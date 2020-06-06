package operator

import (
	"../buffers"
	"../locks"
	"../marshaller"
	"../persist"
	"../protocol"
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
	go ne.keepSendingFromOutgoingMessages()
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
		packet := protocol.Packet{"message", []interface{}{message}}
		serializedPacket := m.Marshall(packet)

		// log.Println(fmt.Sprintf("Sending to conn %s msg %f", subscriber, message))
		locks.OutgoingLock.Lock()
		buffers.OutgoingMessages[subscriber] = append(buffers.OutgoingMessages[subscriber], serializedPacket)
		persist.Save("./database/outgoing", buffers.OutgoingMessages)
		locks.OutgoingLock.Unlock()
	}
}

func (ne *NotificationEngine) keepSendingFromOutgoingMessages() {
	for {
		locks.OutgoingLock.Lock()

		for connectionId, messagesToSend := range buffers.OutgoingMessages {
			numberOfMessagesToSend := len(messagesToSend)

			for i := 0; i < numberOfMessagesToSend; i++ {
				ne.srh.Send(connectionId)
			}
		}

		locks.OutgoingLock.Unlock()
	}
}