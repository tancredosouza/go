package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var mutex = &sync.Mutex{} // mutual-exclusion lock
var bridge string

var isGoingLeft bool
var isGoingRight bool

func keepTryingCrossingBridgeToThe(direction string, n time.Duration) {
	for {
		fmt.Println("🕑", "Car is waiting to go to the", direction)
		time.Sleep(n * time.Second)

		mutex.Lock()

		tryCrossingBridge(direction)

		mutex.Unlock()
	}
}

func tryCrossingBridge(s string) {
	switch s {
	case "left":
		isGoingLeft = true
		cross()
		isGoingLeft = false
		break
	case "right":
		isGoingRight = true
		cross()
		isGoingRight = false
		break
	}
}

func cross() {
	var direction string = getDirection()

	fmt.Println("=================================")
	fmt.Println("🚗", "Car on the bridge is going", direction)

	time.Sleep(5 * time.Second)

	fmt.Println("✅", "Car finished going to the", direction)
	fmt.Println("---------------------------------")
}

func getDirection() string {
	if isGoingLeft {
		return "left"
	} else {
		return "right"
	}
}

func bridgeWatcher() {
	for {
		if isGoingLeft && isGoingRight {
			fmt.Println("💥", "CRASH HAPPENED!")
			os.Exit(-1) // all 1s in binary
		}
	}
}

func main() {
	go keepTryingCrossingBridgeToThe("left", 2)
	go keepTryingCrossingBridgeToThe("right", 3)

	go bridgeWatcher()

	fmt.Scanln()
}
