package cli

import (
	"fmt"

	"github.com/Joshua-Lucas/pokedex-cli/internal/pokeapi"
	"github.com/Joshua-Lucas/pokedex-cli/internal/pokecache"
)

type Config struct {
	Next     string
	Previous string
	Cache    *pokecache.Cache
}

func (c *Config) getNext() error {

	locations, err := pokeapi.GetLocations(c.Next, c.Cache)
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

func (c *Config) getPrev() error {
	onFirstPage := false
	if c.Previous == "" {
		c.Previous = "https://pokeapi.co/api/v2/location-area/"
		onFirstPage = true
	}

	locations, err := pokeapi.GetLocations(c.Previous, c.Cache)
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

	if onFirstPage == true {
		fmt.Println("you're on the first page")
	}

	return nil
}
