package main

import (
	"../component"
	"fmt"
)


func main() {
	for i:=1;i<2;i++ {
		go operate(i)
	}

	fmt.Scanln()
}

func operate(i int) {
	var c = component.Component{Key: i}
	c.TmqConnect("localhost", 3993)

	c.CreateTopic("Sala")
	c.CreateTopic("Quarto 1")
	c.CreateTopic("Quarto 2")
	c.CreateTopic("Cozinha")

	//time.Sleep(5*time.Second)
	c.Publish("Sala", 25)
	for i:=0;i<50;i++{
		c.Publish("Cozinha", 18)
		c.Publish("Sala", 23)
	}
}
