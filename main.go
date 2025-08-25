package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	var cleanedInput []string

	for t := range strings.FieldsSeq(text) {
		cleanedInput = append(cleanedInput, strings.ToLower(t))
	}

	return cleanedInput
}
