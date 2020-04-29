package main

import (
	"../service"
	"fmt"
)

func main() {
	// looks up for stack/queue proxy from remote naming service
	// unmarshalling
	s := service.StackProxy{
		HostIp:         "localhost",
		HostPort:       6966,
		RemoteObjectId: 1,
	}

	// runs operations
	fmt.Println(s.Push(6))
	fmt.Println(s.Push(2))
	fmt.Println(s.Push(3))
	fmt.Println(s.Size())
	fmt.Println(s.Top())
	fmt.Println(s.Pop())
	fmt.Println(s.Size())
	fmt.Println(s.Top())
	fmt.Println(s.Pop())
	fmt.Println(s.Top())
	fmt.Println(s.Pop())
	fmt.Println(s.Size())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
