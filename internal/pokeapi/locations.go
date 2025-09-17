package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Joshua-Lucas/pokedex-cli/internal/pokecache"
)

type pokeWorldLocationsData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type pokeWorldLocationData struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name string
	Url  string
}

func GetLocations(url string, cache *pokecache.Cache) (pokeWorldLocationsData, error) {

	cachedVal, ok := cache.Get(url)

	if ok {
		data := pokeWorldLocationsData{}
		err := json.Unmarshal(cachedVal, &data)

		if err != nil {
			return pokeWorldLocationsData{}, fmt.Errorf("Error unmarshalling location data: %v", err)
		}

		return data, nil

	}

	res, err := http.Get(url)
	if err != nil {
		return pokeWorldLocationsData{}, fmt.Errorf("Error occurred when fetching the location area data")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		return pokeWorldLocationsData{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		return pokeWorldLocationsData{}, fmt.Errorf("read body: %w", err)
	}

	// Store body value in cache
	cache.Add(url, body)

	data := pokeWorldLocationsData{}
	err = json.Unmarshal(body, &data)

	if err != nil {
		return pokeWorldLocationsData{}, fmt.Errorf("Error unmarshalling location data: %v", err)
	}

	return data, nil
}

// Gets the pokemon in the specific location
func GetPokemonInLocation(location string, cache *pokecache.Cache) ([]Pokemon, error) {
	var pokemon []Pokemon

	cachedVal, ok := cache.Get(location)

	// If there is a value in the cache us that value
	if ok {
		data := pokeWorldLocationData{}
		err := json.Unmarshal(cachedVal, &data)

		if err != nil {
			return pokemon, fmt.Errorf("Error unmarshalling location data: %v", err)
		}

		for _, val := range data.PokemonEncounters {
			p := Pokemon{
				Name: val.Pokemon.Name,
				Url:  val.Pokemon.URL,
			}
			pokemon = append(pokemon, p)
		}

		return pokemon, nil

	}

	// If no value in cache was found make request to get pokemon encounters
	url := "https://pokeapi.co/api/v2/location-area/" + location
	res, err := http.Get(url)
	if err != nil {
		return pokemon, fmt.Errorf("Error occurred when fetching the location area data")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		return pokemon, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		return pokemon, fmt.Errorf("read body: %w", err)
	}

	// Store body value in cache
	cache.Add(location, body)

	data := pokeWorldLocationData{}
	err = json.Unmarshal(body, &data)

	if err != nil {
		return pokemon, fmt.Errorf("Error unmarshalling location data: %v", err)
	}

	for _, val := range data.PokemonEncounters {
		p := Pokemon{
			Name: val.Pokemon.Name,
			Url:  val.Pokemon.URL,
		}
		pokemon = append(pokemon, p)
	}

	return pokemon, nil
}
