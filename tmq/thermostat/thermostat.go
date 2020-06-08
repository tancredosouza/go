package main

import (
	"../component"
	"fmt"
	"time"
)


var tmqComponent = InitializeTmqComponent()
func InitializeTmqComponent() component.Component {
	var c = component.Component{Key: 9999}
	c.TmqConnect("localhost", 3993)

	return c
}

func TMQ_ChangeRoomTemperature(room string, desiredTemperature int) {
	tmqComponent.Publish(room, desiredTemperature)
}

func main() {
	for i := 0; i < 10000; i++ {
		//log.Println("Sending temperature...")
		time.Sleep(time.Second)
		TMQ_ChangeRoomTemperature("Sala", 19)
	}
	fmt.Scanln()
}