package main

import (
	"fmt"
	"github.com/my/repo/mymiddleware/constants"
	"github.com/my/repo/mymiddleware/distribution_05"
	"github.com/my/repo/mymiddleware/service"
)

func main() {
	n := service.NamingServiceProxy{NamingServiceIp: "localhost", NamingServicePort: 3999}

	for i:=0; i < 100; i++ {
		register(n, i)
	}

	inv := distribution_05.Invoker{"localhost", 9132}
	inv.Invoke()

	fmt.Scanln()
}

func register(n service.NamingServiceProxy, i int) {
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