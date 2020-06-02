package operator

import (
	"../buffers"
	"../marshaller"
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
	}
}

func createTopic(name string) {
	if _, found := buffers.Topics[name]; !found {
		buffers.Topics[name] = make(chan []byte, 100)
		log.Println(fmt.Sprintf("Created topic with name %s!", name))
	}
}