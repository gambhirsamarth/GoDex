package main

import (
	"GoDex/main/service"
	"GoDex/main/utility"
)

func main() {
	service.InitializeDefaultPokedex()
	pokedex := service.GetPokedex()
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
			pokemon := utility.RequestPokemon()
			utility.AddPokemon(pokemon.Name, pokemon, pokedex)
		case "GET":
			pokemonName := utility.RequestPokemonName()
			utility.GetPokemon(pokemonName, pokedex)
		default:
			utility.PrintInvalidCommandMessage()
		}
	}
	utility.PrintExitMessage()
}
