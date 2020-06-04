package main

import (
	"./buffers"
	"./operator"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	InitializeMiddleware()
}

func InitializeMiddleware() {
	operator.Initialize()
	buffers.Initialize()

	ne := operator.NotificationEngine{ServerHost: "localhost", ServerPort: 3993}
	go ne.Initialize()

	fmt.Scanln()
}
