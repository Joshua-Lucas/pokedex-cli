package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Joshua-Lucas/pokedex-cli/internal/cli"
	"github.com/Joshua-Lucas/pokedex-cli/internal/pokecache"
)

func main() {
	CLI_COMMANDS := map[string]cli.Command{}
	CONFIG := cli.Config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
		Cache:    pokecache.NewCache(5 * time.Second),
	}

	CLI_COMMANDS["exit"] = cli.Command{
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    cli.Exit(&CONFIG),
	}

	CLI_COMMANDS["help"] = cli.Command{
		Name:        "help",
		Description: "Provides usage details to the user",
		Callback:    cli.Help(&CONFIG, CLI_COMMANDS),
	}

	CLI_COMMANDS["map"] = cli.Command{
		Name:        "map",
		Description: "Displays map locations in our pokeman world",
		Callback:    cli.Map(&CONFIG),
	}

	CLI_COMMANDS["mapb"] = cli.Command{
		Name:        "mapb",
		Description: "Display the precious map locations",
		Callback:    cli.MapBack(&CONFIG),
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
				command.Callback()
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
