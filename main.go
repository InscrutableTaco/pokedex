package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	scn := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("Pokedex > ")
		scn.Scan()
		input := scn.Text()

		inputSlice := cleanInput(input)

		if len(inputSlice) > 0 {
			fmt.Println("Your command was:", inputSlice[0])
		}
	}

}
