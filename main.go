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
		fmt.Println(input)
	}
}
