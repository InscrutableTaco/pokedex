package main

import (
	"fmt"
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
	fmt.Print("Hello, World!\n")
}
