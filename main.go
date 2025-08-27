package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	CLI_COMMANDS := map[string]cliCommand{}

	CLI_COMMANDS["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	CLI_COMMANDS["help"] = cliCommand{
		name:        "help",
		description: "Provides usage details to the user",
		callback:    commandHelp(CLI_COMMANDS),
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()

		rawUserInput := scanner.Text()

		cleanedUserInput := cleanInput(rawUserInput)

		for _, val := range cleanedUserInput {

			command, ok := CLI_COMMANDS[val]

			if ok {
				command.callback()
			} else {
				fmt.Println("Unknown command")
			}

		}
	}

}

func cleanInput(text string) []string {
	var cleanedInput []string

	for t := range strings.FieldsSeq(text) {
		cleanedInput = append(cleanedInput, strings.ToLower(t))
	}

	return cleanedInput
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp(commands map[string]cliCommand) func() error {

	return func() error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage: ")
		fmt.Println("")

		for _, val := range commands {
			if val.name == "help" {
				continue
			}

			fmt.Printf("%s: %s\n", val.name, val.description)

		}
		return nil
	}

}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}
