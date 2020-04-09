package main

import (
	"bufio"
	"fmt"
	"net"
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

func main() {
	go startAndRunServer()

	fmt.Scanln()
}

func startAndRunServer() {
	fmt.Println("Server is starting...")
	ln, _ := net.Listen(shared.TCP_PROTOCOL, shared.LOCAL_HOST_IP+":"+shared.DEFAULT_PORT)
	fmt.Println("Server is listening.")

	for {
		clientConnection := listenConnectionOnLocalHost(ln)

		if clientConnection == nil {
			fmt.Println("nulo")
			continue
		}

		go convertDollarCostReadFromConnectionToPokemon(clientConnection)
	}
}

func listenConnectionOnLocalHost(listener net.Listener) net.Conn {
	clientConnection, err := listener.Accept()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return clientConnection
}

func closeConnection(connection net.Conn) {
	fmt.Println("Closing connection...")
	connection.Close()
	fmt.Println("Successfully closed connection!")
}

func convertDollarCostReadFromConnectionToPokemon(clientConnection net.Conn) {
	fmt.Println("Listening to a client")
	for {
		dollarCostInReais, _ := bufio.NewReader(clientConnection).ReadString('\n')
		pokemonNumber := convertMoneyToPokemon(dollarCostInReais)
		validMessageToClient := buildValidMessageToClientFrom(pokemonNumber)
		writeMessageToConnection(validMessageToClient, clientConnection)
	}
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
	if pokemonNumber <= 0 {
		return "Invalid result. Make sure your string matches the format R$xx.xx"
	} else if pokemonNumber >= len(shared.Pokemons) {
		return "... Well.. There's not enough Pokémons."
	} else {
		return "#" + strconv.Itoa(pokemonNumber) + ": " + shared.Pokemons[pokemonNumber-1]
	}
}

func writeMessageToConnection(dollarCostInReais string, connection net.Conn) {
	connection.Write([]byte(dollarCostInReais + "\n"))
}
