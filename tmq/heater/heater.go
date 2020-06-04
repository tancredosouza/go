package main

import (
	"../component"
	"fmt"
	"log"
)

func main() {
	c := component.Component{Key: 101}
	c.TmqConnect("localhost", 3993)

	for {
		log.Println(fmt.Sprintf("Temperatura configurada para %f", <- c.SubscriptionMessages))
	}
}
