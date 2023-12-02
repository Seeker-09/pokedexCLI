package main

import (
	"fmt"
	"github.com/Seeker-09/pokedexCLI/internal/pokeapi"
	"log"
)

func main() {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	//startRepl()
}
