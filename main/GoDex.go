package main

import "fmt"

type Pokemon struct {
	Name    string
	Type    string
	HP      int
	Attack  int
	Defense int
}

var pokedex = map[string]Pokemon{
	"Pikachu": {
		Name:    "Pikachu",
		Type:    "Electric",
		HP:      35,
		Attack:  55,
		Defense: 40,
	},
	"Charmander": {
		Name:    "Charmander",
		Type:    "Fire",
		HP:      39,
		Attack:  52,
		Defense: 43,
	},
	"Bulbasaur": {
		Name:    "Bulbasaur",
		Type:    "Grass",
		HP:      45,
		Attack:  49,
		Defense: 49,
	},
	"Squirtle": {
		Name:    "Squirtle",
		Type:    "Water",
		HP:      44,
		Attack:  48,
		Defense: 65,
	},
	"Eevee": {
		Name:    "Eevee",
		Type:    "Normal",
		HP:      55,
		Attack:  55,
		Defense: 50,
	},
	"Jigglypuff": {
		Name:    "Jigglypuff",
		Type:    "Fairy",
		HP:      115,
		Attack:  45,
		Defense: 20,
	},
	"Meowth": {
		Name:    "Meowth",
		Type:    "Normal",
		HP:      40,
		Attack:  45,
		Defense: 35,
	},
}

func addPokemon(name string, pokemon Pokemon) {
	if _, ok := pokedex[name]; ok {
		return
	} else {
		pokedex[name] = pokemon
	}
}

func getPokemon(name string) {
	if _, ok := pokedex[name]; ok {
		fmt.Println(pokedex[name])
	} else {
		fmt.Println("Pokemon not found")
	}
}

func listAllPokemonNames() {
	for _, pokemon := range pokedex {
		fmt.Println(pokemon.Name)
	}
}

func main() {
	printWelcomeMessage()
	running := true
	for running {
		printAvailableCommands()
		command := requestUserCommand()
		switch command {
		case "EXIT":
			running = false
			break

		case "LIST":
			listAllPokemonNames()
		case "ADD":
			//addPokemon()
			continue // TODO : Complete
		case "GET":
			pokemonName := requestPokemonName()
			getPokemon(pokemonName)
		default:
			printInvalidCommandMessage()
		}
	}
	printExitMessage()
}

func printWelcomeMessage() {
	fmt.Println()
	fmt.Println("Welcome to GoDex!")
}

func printExitMessage() {
	fmt.Println()
	fmt.Println("Thank you for using GoDex!")
}

func printAvailableCommands() {
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println()
	fmt.Println("ADD : to add a Pokemon")
	fmt.Println("LIST : to list all Pokemon names")
	fmt.Println("GET : to get pokemon details")
	fmt.Println("EXIT : to exit GoDex")
}

func printInvalidCommandMessage() {
	fmt.Println()
	fmt.Println("Invalid command")
}

func requestUserCommand() (command string) {
	fmt.Println()
	fmt.Println("Please enter a command: ")
	fmt.Scanln(&command)
	return
}

func requestPokemonName() (pokemonName string) {
	fmt.Println()
	fmt.Println("Please enter the Pokemon name: ")
	fmt.Scanln(&pokemonName)
	return pokemonName
}
