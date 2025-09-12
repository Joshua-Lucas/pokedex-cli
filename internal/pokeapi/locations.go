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
