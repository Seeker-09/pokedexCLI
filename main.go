package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		s.Scan()
		input := s.Text()
		commands := getCommands()
		command := commands[input]
		command.action()
	}
}

type cliCommand struct {
	name        string
	description string
	action      func()
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help Message",
			action:      helpAction,
		},
	}
}

func helpAction() {
	for _, c := range getCommands() {
		fmt.Println("Sure! Here is a list of commands: ")
		fmt.Println("")
		fmt.Printf("%s - %s\n", c.name, c.description)
	}
}
