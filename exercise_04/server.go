package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"regexp"
	"strconv"
)

//RegexCurrencyToNumber ..., defines the regEX used to extract all
//numbers from a money amount (e.g. R$4,70 -> 470)
const RegexCurrencyToNumber = `R\$(?P<Reais>\d+)\,(?P<Cents>\d+)`

//ErrorCode ..., equal to -1, is used as a return value to determine
//something went wrong in the conversion
const ErrorCode = -1

func getPokemonNumberFromDollarAmount(dollarAmount string) int {
	params := getParams(RegexCurrencyToNumber, dollarAmount)

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

func getValidPokemonNameFromPokemonNumber(pokemonNumber int) string {
	if pokemonNumber <= 0 {
		return "Invalid result. Make sure your string matches the format R$xx.xx"
	} else if pokemonNumber >= len(Pokemons) {
		return "... Well.. There's not enough Pokémons."
	} else {
		return "#" + strconv.Itoa(pokemonNumber) + ": " + Pokemons[pokemonNumber-1]
	}
}

func receiveAndConvertMessages(queueName string, msgs (<- chan amqp.Delivery), ch (*amqp.Channel)) {
	for d := range msgs {
		msg := string(d.Body)
		pokemonNumber := getPokemonNumberFromDollarAmount(msg)
		responseMessage := getValidPokemonNameFromPokemonNumber(pokemonNumber)
		PublishMessageToQueue(responseMessage, queueName, ch)
	}
}

func main() {
	conn := CreateConnectionWithHost("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	ch := CreateChannelOnConnection(conn)
	defer ch.Close()

	requestQueue := CreateQueueOnChannel(ch, "request")
	responseQueue := CreateQueueOnChannel(ch, "response")

	msgs := ConsumeFromQueue(requestQueue.Name, ch)

	go receiveAndConvertMessages(responseQueue.Name, msgs, ch)

	fmt.Scanln()
}