package main

import "fmt"

func callBackHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex help menu")
	fmt.Println("Here are your available comands:")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println("")
	return nil
}
