package main

import (
	"fmt"
	"github.com/my/repo/mymiddleware/constants"
	"github.com/my/repo/mymiddleware/distribution"
	"github.com/my/repo/mymiddleware/service"
)

func main() {
	n := service.NamingServiceProxy{NamingServiceIp: "localhost", NamingServicePort: 3999}

	q := service.QueueProxy{HostIp: "localhost", HostPort: 9132, RemoteObjectId: constants.QUEUE_ID, TypeName: "queue"}
	s := service.StackProxy{HostIp:"localhost",HostPort: 9132,RemoteObjectId: constants.STACK_ID, TypeName: "stack"}

	res := n.Register("app.Stack", s)
	fmt.Println(res)

	res = n.Register("app.Queue", q)
	fmt.Println(res)

	inv := distribution.Invoker{"localhost", 9132}
	inv.Invoke()

	fmt.Scanln()
}