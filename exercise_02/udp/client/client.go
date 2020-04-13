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

	for i := 1; i < lim; i++ {
		go startAndRunClient(false, "")
	}
	go startAndRunClient(true, "")

	fmt.Scanln()
}

func startAndRunClient(shouldWriteToFile bool, filename string) {
	serverConnection := startUDPConnectionOnLocalHost()

	outputFile, err := os.Create("times_udp" + strconv.Itoa(1) + ".txt") // creating...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}

	for i := 0; i < 10000; i++ {
		fmt.Println(strconv.Itoa(i))
		reaisAmount := "R$6,12"
		var message []byte = []byte(reaisAmount)

		startTime := time.Now()
		_, err := serverConnection.Write(message)
		if err != nil {
			fmt.Println(err)
		}

		listen(serverConnection)
		endTime := time.Since(startTime)

		if shouldWriteToFile {
			fmt.Fprintln(outputFile, endTime.Seconds())
		}
	}

	serverConnection.Close()
	return
}

func listen(connection *net.UDPConn) {
	buffer := make([]byte, 1024)
	_, _, err := 0, new(net.UDPAddr), error(nil)
	for err == nil {
		n, _, innerError := connection.ReadFromUDP(buffer)

		if innerError != nil {
			fmt.Println(innerError)
			break
		}

		if n != 0 {
		//		fmt.Println(string(buffer[:n]))
			break
		}
	}
}

func generateRandomReaisAmount() string {
	reaisAmount := rand.Intn(10)
	centsAmount := rand.Intn(99)

	return "R$" + strconv.Itoa(reaisAmount) + "," + strconv.Itoa(centsAmount)
}

func startUDPConnectionOnLocalHost() *net.UDPConn {
	RemoteAddr, resolveError := net.ResolveUDPAddr("udp", shared.LOCAL_HOST_IP+":"+shared.DEFAULT_PORT)

	if resolveError != nil {
		fmt.Println(resolveError)
	}

	serverConnection, dialError := net.DialUDP(shared.UDP_PROTOCOL, nil, RemoteAddr)
	if dialError != nil {
		fmt.Println(dialError)
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
	message, _ := bufio.NewReader(serverConnection).ReadString('\n')

	fmt.Print("Pokemon " + message)
}

func writeToFile(t time.Duration, filename string) {
	x := t

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%f\n", x.Seconds()*1000.0)) // writing...
	if err != nil {
		fmt.Printf("error writing string: %v", err)
	}

	fmt.Printf("%f\n", x.Seconds()*1000)
}
