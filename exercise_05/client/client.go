package main

import (
	"../service"
	"fmt"
)

func main() {
	namingProxy := service.NamingServiceProxy{"localhost", 3999}
	queueProxy := namingProxy.Lookup("FilaDoMal")
	stackProxy := namingProxy.Lookup("PilhaDoMal")

	fmt.Println(queueProxy.GetSize())
	fmt.Scanln()
}
