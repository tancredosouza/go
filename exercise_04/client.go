package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

func generateRandomReaisAmount() string {
	rand.Seed(time.Now().UnixNano())
	reaisAmount := rand.Intn(10)
	centsAmount := rand.Intn(99)

	return "R$" + strconv.Itoa(reaisAmount) + "," + strconv.Itoa(centsAmount)
}

func main() {
	conn, _ := amqp.Dial("amqp://127.0.0.1:5672/")
	defer conn.Close()
	fmt.Println("Successfully connected to RabbitMQ.")

	ch, _ := conn.Channel()
	defer ch.Close()
	fmt.Println("Successfully opened channel")

	requestQueue, _ := ch.QueueDeclare(
		"request",
		false,
		false,
		false,
		false,
		nil,
	)

	responseQueue, _ := ch.QueueDeclare(
		"response",
		false,
		false,
		false,
		false,
		nil,
	)

	fmt.Println("Successfully declared queue")

	body := generateRandomReaisAmount()
	fmt.Println(body)
	_ = ch.Publish(
		"",                // exchange
		requestQueue.Name, // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	msgs, _ := ch.Consume(
		responseQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [x] Sent %s", body)

	fmt.Scanln()
}