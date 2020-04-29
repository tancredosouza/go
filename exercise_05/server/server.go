package main

import (
	"../service"
	"fmt"
)

func main() {
	// registers itself on naming service
	n := service.NamingServiceProxy{"localhost", 3245}

	fmt.Println(n.Lookup("FilaDoMal"))
	// starts listening
}