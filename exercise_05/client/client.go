package main

import (
	"../service"
	"fmt"
)

func main() {
	namingProxy := service.NamingServiceProxy{"localhost", 3999}
	queueProxy := namingProxy.Lookup("FilaDoMal")

	fmt.Println(queueProxy.GetSize())
	fmt.Scanln()
}
