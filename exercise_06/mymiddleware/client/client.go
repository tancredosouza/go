package main

import (
	"fmt"
	"github.com/my/repo/mymiddleware/service"
	"math/rand"
	"time"
)

func main() {
	//outputFile, _ := os.Create("timemymiddleware_seconds.txt")

	namingProxy := service.NamingServiceProxy{NamingServiceIp:"localhost", NamingServicePort:3999}

	for i := 0; i < 100; i++ {
		queueProxy := namingProxy.Lookup(fmt.Sprintf("app.Queue_%d", i))
		fmt.Println(i)

		go performRandomOperations(queueProxy.(service.QueueProxy))
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

func performRandomOperations(queueProxy service.QueueProxy) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 500; i++ {
		r := rand.Intn(4)

		if (r == 0) {
			x := rand.Intn(100)
			queueProxy.InsertElement(x)
		}

		if (r == 1) {
			fmt.Println(queueProxy.GetSize())
		}

		if (r == 2) {
			fmt.Println(queueProxy.RemoveElement())
		}

		if (r == 3) {
			fmt.Println(queueProxy.GetFirstElement())
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