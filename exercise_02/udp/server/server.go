package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"

	"../../shared"
)

//RegexCurrencyToNumber ..., defines the regEX used to extract all
//numbers from a money amount (e.g. R$4,70 -> 470)
const RegexCurrencyToNumber = `R\$(?P<Reais>\d+)\,(?P<Cents>\d+)`

//ErrorCode ..., equal to -1, is used as a return value to determine
//something went wrong in the conversion
const ErrorCode = -1

const NUMBER_OF_CONNECTIONS = 5

func main() {
	startAndRunServer()

	fmt.Scanln()
}

func startAndRunServer() {
	conn := startServer()
	runServer(conn)
}

func startServer() *net.UDPConn {
	fmt.Println("Server is starting...")
	udpAddress, resolveError := net.ResolveUDPAddr("udp", shared.LOCAL_HOST_IP+":"+shared.DEFAULT_PORT)
	if resolveError != nil {
		fmt.Println(resolveError)
		os.Exit(0)
	}

	conn, listeningError := net.ListenUDP("udp", udpAddress)

	if listeningError != nil {
		fmt.Println(listeningError)
		os.Exit(0)
	}

	fmt.Println("Server successfully created an UDP connection.")
	return conn
}

func runServer(conn *net.UDPConn) {
	fmt.Println("Server will now start " + strconv.Itoa(NUMBER_OF_CONNECTIONS) + " connections.")

	for i := 0; i < NUMBER_OF_CONNECTIONS; i++ {
		go keepListeningToUDPConnection(conn)
	}

	fmt.Println("Server is ready.")
}

func keepListeningToUDPConnection(connection *net.UDPConn) {
	userMessage := make([]byte, 1024)

	for {
		n, addr, readingError := connection.ReadFromUDP(userMessage)

		if readingError != nil {
			fmt.Println(readingError)
			break
		}

		dollarCostInReais := string(userMessage[:n])
		pokemonName := convertDollarCostReadFromConnectionToPokemon(dollarCostInReais)

		connection.WriteToUDP([]byte(pokemonName), addr)
	}

	defer closeConnection(connection)
}

func closeConnection(connection net.Conn) {
	fmt.Println("Closing connection...")
	connection.Close()
	fmt.Println("Successfully closed connection!")
}

func convertDollarCostReadFromConnectionToPokemon(dollarCostInReais string) string {

	pokemonNumber := convertMoneyToPokemon(dollarCostInReais)
	validMessageToClient := buildValidMessageToClientFrom(pokemonNumber)

	return validMessageToClient

}

func convertMoneyToPokemon(dollarCostInReais string) int {
	params := getParams(RegexCurrencyToNumber, dollarCostInReais)

	if len(params) == 0 {
		return ErrorCode
	}

	moneyAmountAsText := params["Reais"] + params["Cents"]
	pokemonNumber, err := strconv.Atoi(moneyAmountAsText)

	if err != nil {
		fmt.Println(err)
		return ErrorCode
	}

	return pokemonNumber
}

//Parses url with the given regular expression and
//returns the group values defined in the expression.
func getParams(regEx, url string) (paramsMap map[string]string) {
	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(url)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return
}

func buildValidMessageToClientFrom(pokemonNumber int) string {
	if pokemonNumber == ErrorCode || pokemonNumber == 0 {
		return "Invalid result."
	} else if pokemonNumber >= len(shared.Pokemons) {
		return "... Well.. There's not enough Pokémons."
	} else {
		return "#" + strconv.Itoa(pokemonNumber) + ": " + shared.Pokemons[pokemonNumber-1]
	}
}
