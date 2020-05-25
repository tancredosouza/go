package main

import (
	"fmt"
	"github.com/my/repo/mymiddleware/result_callback"
	"github.com/my/repo/mymiddleware/service"
	"log"
	"os"
)

var outputFile, _ = os.Create("time_100clients_prevexercise.txt")

var rcvMsgs chan string

func main() {
	rcvMsgs = make(chan string, 1)
	namingProxy := service.NamingServiceProxy{HostIp: "localhost", HostPort:3999}

	for i := 0; i < 1; i++ {
		namingProxy.Initialize()
		queueProxy := namingProxy.Lookup(fmt.Sprintf("app.Queue_%d", i))
		go performOperations(false, *queueProxy)
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

func performOperations(write bool, queueProxy service.QueueProxy) {
	r := result_callback.ResultCallback{}
	r.Initialize()

	queueProxy.Initialize(&r)
	queueProxy.InsertElement(33)
	i := 0
	rcv := 0
	for {
		select {
			case ans := <- result_callback.ReceivedMsgs:
				log.Println(rcv, "--->", ans)
				rcv++
			default:
				if (i < 10000) {
					queueProxy.GetFirstElement()
					i++
				} else if (rcv >= 10000){
					break
				}
		}
	}
}