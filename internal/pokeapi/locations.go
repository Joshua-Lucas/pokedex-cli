package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func GetLocations(url string) (pokeWorldLocationsData, error) {

	res, err := http.Get(url)
	if err != nil {
		return pokeWorldLocationsData{}, fmt.Errorf("Error occurred when fetching the location area data")
	}

	body, err := io.ReadAll(res.Body)

	res.Body.Close()

	if res.StatusCode > 299 {
		return pokeWorldLocationsData{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		return pokeWorldLocationsData{}, fmt.Errorf("Error occurred")
	}

	data := pokeWorldLocationsData{}
	err = json.Unmarshal(body, &data)

	if err != nil {
		return pokeWorldLocationsData{}, fmt.Errorf("Error unmarshalling location data: %v", err)
	}

	return data, nil
}
