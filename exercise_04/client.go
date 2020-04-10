package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func generateRandomReaisAmount() string {
	rand.Seed(time.Now().UnixNano())
	reaisAmount := rand.Intn(10)
	centsAmount := rand.Intn(99)

	return "R$" + strconv.Itoa(reaisAmount) + "," + strconv.Itoa(centsAmount)
}

func keepReceivingMessagesFromChannel(msgs (<- chan amqp.Delivery)) {
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
	}
}

func main() {
	conn := CreateConnectionWithHost("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	ch := CreateChannelOnConnection(conn)
	defer ch.Close()

	requestQueue  := CreateQueueOnChannel(ch, "request")
	responseQueue := CreateQueueOnChannel(ch, "response")

	msgs := ConsumeFromQueue(responseQueue.Name, ch)

	for i := 0; i < 10000; i++ {
		reaisAmount := generateRandomReaisAmount()
		PublishMessageToQueue(reaisAmount, requestQueue.Name, ch)
		//a := <- msgs

		//log.Printf(strconv.Itoa(i) + "  -- " + string(a.Body))
	}

	go keepReceivingMessagesFromChannel(msgs)

	fmt.Scanln()
}