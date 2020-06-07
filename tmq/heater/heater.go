package main

import (
	"github.com/streadway/amqp"
	"log"
)
/*
var tmqComponent = InitializeTmqComponent()
func InitializeTmqComponent() component.Component {
	var c = component.Component{Key: 101}
	c.TmqConnect("localhost", 3993)

	return c
}

func TMQ_ListenForChangesInTemperature(room string) {
	go func() {
		for {
			fmt.Println("Temp received = ", <- tmqComponent.SubscriptionMessages)
		}
	}()

	fmt.Scanln()
}
*/

var rabbitMqChannel = RabbitMQ_Initialize()
func RabbitMQ_Initialize() (*amqp.Channel) {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()

	return ch
}

func RabbitMQ_ListenForTemperatureChange(room string) {
	forever := make(chan bool)

	msgs, _ := rabbitMqChannel.Consume(
		room, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	go func() {
		for {
			for d := range msgs {
				log.Printf("Received a message: %s", d.Body)
			}
		}
	}()

	<- forever
}

func main() {
	// TMQ_ListenForChangesInTemperature("Sala")
	RabbitMQ_ListenForTemperatureChange("Sala")
}