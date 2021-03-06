package main

import (
	"fmt"
	"github.com/my/repo/mymiddleware/constants"
	"github.com/my/repo/mymiddleware/distribution"
	"github.com/my/repo/mymiddleware/service"
)

func main() {
	n := service.NamingServiceProxy{HostIp: "localhost", HostPort: 3999}
	n.Initialize()

	for i:=0; i < 2; i++ {
		register(n, i)
	}

	inv := distribution.Invoker{"localhost", 9132}
	inv.Invoke()

	fmt.Scanln()
}

func register(n service.NamingServiceProxy, i int) {
	q := service.QueueProxy{
		HostIp: "localhost",
		HostPort: 9132,
		RemoteObjectId: constants.QUEUE_ID,
		TypeName: "queue",
		QueueNumber: i}

	res := n.Register(fmt.Sprintf("app.Queue_%d", i), q)

	fmt.Println(res, " ", i)
}