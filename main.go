package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commandMap map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println("")
	for _, cmd := range commandMap {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
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
	}

	scn := bufio.NewScanner(os.Stdin)

	for true {
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
