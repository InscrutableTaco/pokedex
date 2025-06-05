package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"sort"
	"strings"
)

const BASE_PATH = "https://pokeapi.co/api/v2/"

var mapOfCommands map[string]cliCommand
var cnfg *config

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	next     string
	previous string
}

func updatePagination(cnfg *config, results pokeapi.Response) {
	if results.Next != nil {
		cnfg.next = *results.Next
	} else {
		cnfg.next = ""
	}
	if results.Previous != nil {
		cnfg.previous = *results.Previous
	} else {
		cnfg.previous = ""
	}
}

func commandHelp(_ *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println("")

	//create a slice of cliCommands
	var cmdSlice []cliCommand
	for _, cmd := range mapOfCommands {
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

func commandExit(_ *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cnfg *config) error {

	var url string
	if cnfg.next != "" {
		url = cnfg.next
	} else {
		url = BASE_PATH + "location-area/"
	}

	results, err := pokeapi.GetAndParse(url)

	if err != nil {
		return err
	}

	updatePagination(cnfg, results)

	for _, loc := range results.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cnfg *config) error {

	var url string
	if cnfg.previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	url = cnfg.previous

	results, err := pokeapi.GetAndParse(url)

	if err != nil {
		return err
	}

	updatePagination(cnfg, results)

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

	cnfg = &config{
		next:     "",
		previous: "",
	}

	mapOfCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback: func() error {
				return commandExit(cnfg)
			},
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback: func() error {
				return commandHelp(cnfg)
			},
		},
		"map": {
			name:        "map",
			description: "Displays the next page of location areas",
			callback: func() error {
				return commandMap(cnfg)
			},
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of location areas",
			callback: func() error {
				return commandMapb(cnfg)
			},
		},
	}

	scn := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scn.Scan()
		input := scn.Text()

		inputSlice := cleanInput(input)

		cmd, ok := mapOfCommands[inputSlice[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		cmd.callback()
	}

}
