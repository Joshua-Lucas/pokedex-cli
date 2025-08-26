package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()

		rawUserInput := scanner.Text()

		cleanedUserInput := cleanInput(rawUserInput)

		fmt.Printf("Your command was: %s\n", cleanedUserInput[0])
	}

}

func cleanInput(text string) []string {
	var cleanedInput []string

	for t := range strings.FieldsSeq(text) {
		cleanedInput = append(cleanedInput, strings.ToLower(t))
	}

	return cleanedInput
}
