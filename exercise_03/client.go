package main

import (
	"bufio"
	"exercise_03/converters"
	shared2 "exercise_03/shared"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {

	lim := 1
	_, err := os.Create("time_in_seconds" + strconv.Itoa(lim) + ".txt") // creating...
	if err != nil {
			fmt.Printf("error creating file: %v", err)
			return
		}
	for i := 1; i < lim; i++ {
		go startAndRunClient(false, "")
	}
	go startAndRunClient(true, "time_in_seconds"+strconv.Itoa(lim)+".txt")

	fmt.Scanln()
}

func startAndRunClient(shouldWrite bool, filename string) {
	serverConnection := startGRPCConnectionOnLocalHost()

	converter := converters.NewGreeterClient(serverConnection)
	ctx := context.Background()

	for i := 0; i < 10000; i++ {
		x := "R$6,12"
		startTime := time.Now()

		_, err := converter.ConvertToPokemon(ctx, &converters.Request{DollarInReais: x})

		if err != nil {
			fmt.Println(err)
		} //else {
		//	fmt.Println(strconv.Itoa(i) + ": " + response.PokemonName)
		//}

		endTime := time.Since(startTime)

		if shouldWrite {
			writeToFile(endTime, filename)
		}
	}
	serverConnection.Close()
}
func generateRandomReaisAmount() string {
	reaisAmount := rand.Intn(10)
	centsAmount := rand.Intn(99)

	return "R$" + strconv.Itoa(reaisAmount) + "," + strconv.Itoa(centsAmount)
}

func startGRPCConnectionOnLocalHost() *grpc.ClientConn {
	serverConnection, err :=
		grpc.Dial(shared2.LOCAL_HOST_IP+":"+shared2.DEFAULT_PORT, grpc.WithInsecure())

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
