package main

import (
	"fmt"
	"net/rpc"
	"os"
)

func main() {
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

	err = client.Call("app.Stack.GetFirstElement", 2, &reply)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(reply)
}