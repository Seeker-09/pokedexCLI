package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Config struct {
	nextUrl     string
	previousUrl string
}

type Locations struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
}

func GetLocationAreas(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)

	err = res.Body.Close()
	if err != nil {
		return
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	printLocationAreas(&body)
}

func printLocationAreas(body *[]byte) {
	var locations Locations
	err := json.Unmarshal(*body, &locations)
	if err != nil {
		fmt.Printf("Error %s\n", err)
		return
	}
	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}
}
