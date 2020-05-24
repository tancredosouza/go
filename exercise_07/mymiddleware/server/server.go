package main

import (
	"fmt"
	"github.com/my/repo/mymiddleware/constants"
	"github.com/my/repo/mymiddleware/distribution"
	"github.com/my/repo/mymiddleware/service"
)

func main() {
	for i:=0; i < 100; i++ {
		register(i)
	}

	inv := distribution.Invoker{"localhost", 9132}
	inv.Invoke()

	fmt.Scanln()
}

func register(i int) {
	n := service.NamingServiceProxy{HostIp: "localhost", HostPort: 3999}
	n.Initialize()

	q := service.QueueProxy{
		HostIp: "localhost",
		HostPort: 9132,
		RemoteObjectId: constants.QUEUE_ID,
		TypeName: "queue",
		QueueNumber: i}

	res := n.Register(fmt.Sprintf("app.Queue_%d", i), q)

	fmt.Println(res, " ", i)
}