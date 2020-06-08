package main

import (
	"../component"
	"fmt"
)

var tmqComponent = InitializeTmqComponent()
func InitializeTmqComponent() component.Component {
	var c = component.Component{Key: 101}
	c.TmqConnect("localhost", 3993)

	return c
}

func TMQ_ListenForChangesInTemperature(room string) {
	for {
		fmt.Println("Temp received = ", <- tmqComponent.SubscriptionMessages)
	}

	fmt.Scanln()
}

func main() {
	tmqComponent.Subscribe("Sala")
	TMQ_ListenForChangesInTemperature("Sala")
}