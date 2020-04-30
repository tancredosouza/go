package main

import (
	"../constants"
	"../distribution"
	"../service"
	"fmt"
)

func main() {
	n := service.NamingServiceProxy{"localhost", 3999}

	q := service.QueueProxy{"localhost", 9132, constants.QUEUE_ID, "queue"}
	s := service.StackProxy{"localhost", 9132, constants.STACK_ID, "stack"}

	queue := n.Register("FilaDoMal", q)
	stk := n.Register("PilhaDoMal", s)

	fmt.Println(queue)
	fmt.Println(stk)

	inv := distribution.Invoker{"localhost", 9132}
	inv.Invoke()

	fmt.Scanln()
}