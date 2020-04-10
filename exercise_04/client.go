package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func randomAmountOfReais() string {
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

	responseMessages := ConsumeFromQueue(responseQueue.Name, ch)

	//go keepReceivingMessagesFromChannel(responseMessages)

	sampleSize := 10000
	for i := 0; i < sampleSize; i++ {
		reaisAmount := randomAmountOfReais()

		//timeStart := time.Now()

		PublishMessageToQueue(reaisAmount, requestQueue.Name, ch)

		<-responseMessages

		//timeEnd := time.Since(timeStart)
		//log.Print(timeEnd.Seconds())
	}


	fmt.Scanln()
}