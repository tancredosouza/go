package main

import (
	"../service"
	"fmt"
)

func main() {
	namingProxy := service.NamingServiceProxy{"localhost", 3999}
	queueProxy := namingProxy.Lookup("app.Queue")
	stackProxy := namingProxy.Lookup("app.Stack")

	test(queueProxy)
	test(stackProxy)
}

func test(proxy service.Proxy) {
	fmt.Println(proxy.InsertElement(5))
	fmt.Println(proxy.InsertElement(2))
	fmt.Println(proxy.InsertElement(4))
	fmt.Println(proxy.InsertElement(6))
	fmt.Println(proxy.InsertElement(1))

	fmt.Println(proxy.GetSize())
	fmt.Println(proxy.RemoveElement())
	fmt.Println(proxy.RemoveElement())
	fmt.Println(proxy.GetSize())
	fmt.Println(proxy.GetFirstElement())

	fmt.Println("---------------------------------")
}