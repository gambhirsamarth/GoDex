package utility

import "fmt"

func RequestUserCommand() (command string) {
	fmt.Println()
	fmt.Println("Please enter a command: ")
	fmt.Scanln(&command)
	return
}

func RequestPokemonName() (pokemonName string) {
	fmt.Println()
	fmt.Println("Please enter the Pokemon name: ")
	fmt.Scanln(&pokemonName)
	return pokemonName
}
