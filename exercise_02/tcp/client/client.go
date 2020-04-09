package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	"../../shared"
)

func main() {

	lim := 1
	//_, err := os.Create("times_tcp" + strconv.Itoa(lim) + ".txt") // creating...
	//if err != nil {
	//		fmt.Printf("error creating file: %v", err)
	//		return
	//	}

	for i := 1; i < lim; i++ {
		go startAndRunClient(false, "")
	}
	go startAndRunClient(false, "")

	fmt.Scanln()
}

func startAndRunClient(shouldWrite bool, filename string) {
	serverConnection := startTCPConnectionOnLocalHost()
	for i := 0; i < 10000; i++ {
		//startTime := time.Now()

		runClientWithConnection(serverConnection)

		//endTime := time.Since(startTime)

		//if shouldWrite {
		//	writeToFile(endTime, filename)
		//}
	}
	serverConnection.Close()
}

func runClientWithConnection(serverConnection net.Conn) {
	dollarCost := generateRandomReaisAmount()

	fmt.Fprintf(serverConnection, dollarCost+"\n")

	getAndPrintPokemonFrom(serverConnection)
}

func generateRandomReaisAmount() string {
	reaisAmount := rand.Intn(10)
	centsAmount := rand.Intn(99)

	return "R$" + strconv.Itoa(reaisAmount) + "," + strconv.Itoa(centsAmount)
}

func startTCPConnectionOnLocalHost() net.Conn {
	serverConnection, err := net.Dial(shared.TCP_PROTOCOL, shared.LOCAL_HOST_IP+":"+shared.DEFAULT_PORT)

	if err != nil {
		fmt.Println(err)
	}

	return serverConnection
}

func readDollarCostInReaisFromUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How much is 1 USD in BRL? ")
	dollarCost, _ := reader.ReadString('\n')

	return dollarCost
}

func getAndPrintPokemonFrom(serverConnection net.Conn) {
	message, err := bufio.NewReader(serverConnection).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Pokemon " + message)
}

func writeToFile(t time.Duration, filename string) {
	x := t

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Error opening file.")
		os.Exit(2)
	}

	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%f\n", x.Seconds()*1000.0)) // writing...
	if err != nil {
		fmt.Printf("error writing string: %v", err)
	}

	fmt.Printf("%f\n", x.Seconds()*1000.0)
}
