package main

import (
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"time"
)

/*
var tmqComponent = InitializeTmqComponent()
func InitializeTmqComponent() component.Component {
	var c = component.Component{Key: 9999}
	c.TmqConnect("localhost", 3993)

	return c
}

func TMQ_ChangeRoomTemperature(room string, desiredTemperature int) {
	tmqComponent.Publish(room, desiredTemperature)
}
*/

var rabbitMqChannel = RabbitMQ_Initialize()
func RabbitMQ_Initialize() (*amqp.Channel) {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()

	return ch
}

func RabbitMQ_ChangeRoomTemperature(room string, desiredTemperature int) {
	q, _ := rabbitMqChannel.QueueDeclare(
		room, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	body := strconv.Itoa(desiredTemperature)
	_ = rabbitMqChannel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(body),
		})
}

func main() {
	for i := 0; i < 10000; i++ {
		log.Println("Sending temperature...")
		time.Sleep(time.Second)
		//TMQ_ChangeRoomTemperature("Sala", 19)
		RabbitMQ_ChangeRoomTemperature("Sala", 19)
	}
}