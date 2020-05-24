package main

import (
	"fmt"
	"github.com/my/repo/mymiddleware/service"
	"log"
	"os"
	"time"
)

var outputFile, _ = os.Create("time_100clients_prevexercise.txt")

func main() {
	namingProxy := service.NamingServiceProxy{HostIp: "localhost", HostPort:3999}

	for i := 0; i < 100; i++ {
		namingProxy.Initialize()
		queueProxy := namingProxy.Lookup(fmt.Sprintf("app.Queue_%d", i))
		go performOperations(i == 0, queueProxy)
	}

	fmt.Scanln()
	return;

	/*
		for i:=0; i<10000;i++ {
			st := time.Now()
			queueProxy.GetFirstElement()

			end := time.Since(st)
			fmt.Fprintln(outputFile, end.Seconds())
			fmt.Println(i)
			t := float64(time.Second) * 0.01
			time.Sleep(time.Duration(t))
		}
	*/
}

func performOperations(write bool, queueProxy *service.QueueProxy) {
	time.Sleep(time.Second)

	queueProxy.Initialize()
	log.Println("------------->", queueProxy.InsertElement(33))

	for x :=0; x < 10000; x++ {
		st := time.Now()
		queueProxy.GetFirstElement()
		end := time.Since(st)

		if write {
			fmt.Fprintln(outputFile, end.Seconds())
			fmt.Println(x)
		}
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