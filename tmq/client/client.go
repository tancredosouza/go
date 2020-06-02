package main

import (
	"../component"
	"log"
)

func main() {
	c := component.Component{}
	c.Dial("localhost", 3993)

	c.Publish("Olar")
}
