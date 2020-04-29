package main

import (
	"../service"
	"fmt"
)

func main() {
	s := service.StackProxy{
		HostIp: "localhost",
		Port: 6966,
		RemoteObjectId: 1,
	}

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
