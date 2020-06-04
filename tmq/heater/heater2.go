package main

import (
	"../component"
	"fmt"
	"log"
)

func main() {
	c := component.Component{Key: 820}
	c.TmqConnect("localhost", 3993)

	c.Subscribe("Sala")
	c.Subscribe("Sala")
	c.Subscribe("Sala")
	c.Subscribe("Sala")
	c.Subscribe("Sala")
	c.Subscribe("Sala")

	for {
		msg := <- c.SubscriptionMessages
		log.Println(fmt.Sprintf("Temperatura configurada para %f", msg.Params[0].(float64)))
	}
}
