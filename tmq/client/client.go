package main

import (
	"../component"
)

func main() {
	c := component.Component{Key: 1}
	c.Dial("localhost", 3993)
}
