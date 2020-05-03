package main

import (
	"fmt"
	"net/rpc"
	"os"
	"time"
)

func main() {
	outputFile, err := os.Create("timegorpc_seconds.txt")

	client, err := rpc.Dial("tcp", "localhost:1567")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	var reply string
	err = client.Call("app.Stack.InsertElement", 33, &reply)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	for i:=0;i < 10000; i++ {
		st := time.Now()
		err = client.Call("app.Stack.GetFirstElement", 2, &reply)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		end := time.Since(st)
		fmt.Fprintln(outputFile, end.Seconds())
	}

	fmt.Println(reply)
}