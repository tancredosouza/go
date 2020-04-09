package main

import (
	"exercise_04/shared"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"regexp"
	"strconv"
)

//RegexCurrencyToNumber ..., defines the regEX used to extract all
//numbers from a money amount (e.g. R$4,70 -> 470)
const RegexCurrencyToNumber = `R\$(?P<Reais>\d+)\,(?P<Cents>\d+)`

//ErrorCode ..., equal to -1, is used as a return value to determine
//something went wrong in the conversion
const ErrorCode = -1

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

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
	} else if pokemonNumber >= len(shared.Pokemons) {
		return "... Well.. There's not enough Pokémons."
	} else {
		return "#" + strconv.Itoa(pokemonNumber) + ": " + shared.Pokemons[pokemonNumber-1]
	}
}


func main() {
	conn, err := amqp.Dial("amqp://127.0.0.1:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	fmt.Println("Successfully connected to RabbitMQ.")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	fmt.Println("Successfully opened channel")

	requestQueue, err := ch.QueueDeclare(
		"request",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	responseQueue, err := ch.QueueDeclare(
		"response",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	fmt.Println("Successfully declared queue")

	msgs, _ := ch.Consume(
		requestQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	go func() {
		for d := range msgs {
			responseMessage := getValidPokemonNameFromPokemonNumber(getPokemonNumberFromDollarAmount(string(d.Body)))

			err = ch.Publish(
				"",
				responseQueue.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(responseMessage),
				})

			failOnError(err, "Failed to publish message.")
		}
	}()

	failOnError(err, "Failed to publish a message")

	fmt.Scanln()
}