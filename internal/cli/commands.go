package cli

import (
	"fmt"
	"os"
)

type Command struct {
	Name        string
	Description string
	Callback    func() error
}

func Help(cfg *Config, commands map[string]Command) func() error {

	return func() error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage: ")
		fmt.Println("")

		for _, val := range commands {
			if val.Name == "help" {
				continue
			}

			fmt.Printf("%s: %s\n", val.Name, val.Description)

		}
		return nil
	}

}

func Exit(cfg *Config) func() error {

	return func() error {
		fmt.Println("Closing the Pokedex... Goodbye!")
		os.Exit(0)

		return nil
	}
}

func Map(cfg *Config) func() error {

	return func() error {

		err := cfg.getNext()
		if err != nil {
			return fmt.Errorf("Error getting next locations: %v", err)
		}

		return nil
	}
}

func MapBack(cfg *Config) func() error {

	return func() error {

		err := cfg.getPrev()
		if err != nil {
			return fmt.Errorf("Error getting previous locations locations: %v", err)
		}

		return nil
	}
}
