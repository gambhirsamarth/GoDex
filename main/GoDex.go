package main

import (
	"GoDex/main/model"
	"GoDex/main/utility"
)

var pokedex = map[string]model.Pokemon{
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

func main() {
	utility.PrintWelcomeMessage()
	running := true
	for running {
		utility.PrintAvailableCommands()
		command := utility.RequestUserCommand()
		switch command {
		case "EXIT":
			running = false
			break

		case "LIST":
			utility.ListAllPokemonNames(pokedex)
		case "ADD":
			//addPokemon()
			continue // TODO : Complete
		case "GET":
			pokemonName := utility.RequestPokemonName()
			utility.GetPokemon(pokemonName, pokedex)
		default:
			utility.PrintInvalidCommandMessage()
		}
	}
	utility.PrintWelcomeMessage()
}
