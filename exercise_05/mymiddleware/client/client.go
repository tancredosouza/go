package main

import (
	"fmt"
	"github.com/my/repo/mymiddleware/service"
	"os"
	"time"
)

func main() {
	outputFile, _ := os.Create("timemymiddleware_seconds.txt")

	namingProxy := service.NamingServiceProxy{NamingServiceIp:"localhost", NamingServicePort:3999}
	stackProxy := namingProxy.Lookup("app.Stack")

	fmt.Println(stackProxy.InsertElement(33))

	for i:=0; i<10000;i++ {
		st := time.Now()
		stackProxy.GetFirstElement()

		end := time.Since(st)
		fmt.Fprintln(outputFile, end.Seconds())
		fmt.Println(i)
		t := float64(time.Second) * 0.01
		time.Sleep(time.Duration(t))
	}
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