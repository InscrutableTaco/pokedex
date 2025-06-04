package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
)

const BASE_PATH = "https://pokeapi.co/api/v2/"

var commandMap map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	next     string
	previous string
}

type location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type response struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []location `json:"results"`
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println("")

	//create a slice of cliCommands
	var cmdSlice []cliCommand
	for _, cmd := range commandMap {
		cmdSlice = append(cmdSlice, cmd)
	}

	//sort it alphabetically
	sort.Slice(cmdSlice, func(i, j int) bool {
		return cmdSlice[i].name < cmdSlice[j].name
	})

	//print each command & description
	for _, cmd := range cmdSlice {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMapfunc() error {

	res, err := http.Get(BASE_PATH + "location-area/")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var results response
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&results); err != nil {
		fmt.Println(err)
		return err
	}

	for _, loc := range results.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func cleanInput(text string) []string {
	var output []string

	splits := strings.Split(text, " ")

	for _, word := range splits {
		trimmedWord := strings.Trim(word, " ")
		if trimmedWord != "" {
			output = append(output, strings.ToLower(trimmedWord))
		}
	}

	return output
}

func main() {

	commandMap = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the location areas",
			callback:    commandMapfunc,
		},
	}

	scn := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scn.Scan()
		input := scn.Text()

		inputSlice := cleanInput(input)

		cmd, ok := commandMap[inputSlice[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		cmd.callback()
	}

}
