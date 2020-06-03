package main

import (
	"../component"
)

func main() {
	c := component.Component{Key: 1}
	c.TmqConnect("localhost", 3993)

	c.CreateTopic("Sala")
	c.CreateTopic("Quarto 1")
	c.CreateTopic("Quarto 2")
	c.CreateTopic("Cozinha")
}
