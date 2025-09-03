package main

import (
	"bufio"
	"fmt"
	"github.com/Joshua-Lucas/pokedex-cli/internal/pokeapi"
	"os"
	"strings"
)

func main() {
	CLI_COMMANDS := map[string]cliCommand{}
	CONFIG := config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
	}

	CLI_COMMANDS["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit(&CONFIG),
	}

	CLI_COMMANDS["help"] = cliCommand{
		name:        "help",
		description: "Provides usage details to the user",
		callback:    commandHelp(&CONFIG, CLI_COMMANDS),
	}

	CLI_COMMANDS["map"] = cliCommand{
		name:        "map",
		description: "Displays map locations in our pokeman world",
		callback:    commandMap(&CONFIG),
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

type config struct {
	Next     string
	Previous string
}

func (c *config) getNext() error {
	locations, err := pokeapi.GetLocations(c.Next)
	if err != nil {
		return fmt.Errorf("Error occurred when fetching map locations: %v", err)
	}

	// Set next config
	c.Next = locations.Next
	c.Previous = locations.Previous

	// Print content
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandHelp(cfg *config, commands map[string]cliCommand) func() error {

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

func commandExit(cfg *config) func() error {

	return func() error {
		fmt.Println("Closing the Pokedex... Goodbye!")
		os.Exit(0)

		return nil
	}
}

func commandMap(cfg *config) func() error {

	return func() error {

		err := cfg.getNext()
		if err != nil {
			return fmt.Errorf("Error getting next locations: %v", err)
		}

		return nil
	}
}
