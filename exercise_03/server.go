package main

import (
	"context"
	"exercise_03/converters"
	shared "exercise_03/shared"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"regexp"
	"strconv"
)

type greeterServer struct {}

//RegexCurrencyToNumber ..., defines the regEX used to extract all
//numbers from a money amount (e.g. R$4,70 -> 470)
const RegexCurrencyToNumber = `R\$(?P<Reais>\d+)\,(?P<Cents>\d+)`

//ErrorCode ..., equal to -1, is used as a return value to determine
//something went wrong in the conversion
const ErrorCode = -1

func (c *greeterServer) ConvertToPokemon(ctx context.Context, in *converters.Request) (*converters.Reply, error) {
	var pokemonNumber = getPokemonNumberFromDollarAmount(in.DollarInReais)

	return &converters.Reply{PokemonName: getValidPokemonNameFromPokemonNumber(pokemonNumber)}, nil
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
	conn, err := net.Listen("tcp", ":"+shared.DEFAULT_PORT)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	servidor := grpc.NewServer()
	converters.RegisterGreeterServer(servidor, &greeterServer{})

	fmt.Println("Servidor pronto ...")

	// Register reflection service on gRPC servidor.
	reflection.Register(servidor)

	err = servidor.Serve(conn);
	if (err != nil) {
		fmt.Println(err)
	}
}
