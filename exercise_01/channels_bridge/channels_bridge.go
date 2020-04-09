package main

import (
	"fmt"
	"time"
)

func finishCrossingBridge(bridge chan string) {
	for {
		value := <-bridge
		fmt.Println("=================================")
		fmt.Println("🚗", "Car on the bridge is going", value)
		fmt.Println("---------------------------------")
		time.Sleep(2 * time.Second)
	}
}

func startCrossingBridge(bridge chan string, v string) {
	for {
		fmt.Println("🕑", "Car is waiting to go to the", v)
		bridge <- v
		time.Sleep(5 * time.Second)
	}
}

func main() {
	bridge := make(chan string)

	go startCrossingBridge(bridge, "left")
	go startCrossingBridge(bridge, "right")
	go finishCrossingBridge(bridge)

	fmt.Scanln()
}
