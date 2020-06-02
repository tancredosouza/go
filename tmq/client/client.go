package main

import (
	"../component"
)

func main() {
	c := component.Component{}
	c.Dial("localhost", 3993)

	c.Publish("Olar")
}
