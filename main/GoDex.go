package main

import (
	"GoDex/main/service"
	"GoDex/main/utility"
	"strings"
)

func main() {
	service.InitializeDefaultPokedex()
	pokedex := service.GetPokedex()
	utility.PrintWelcomeMessage()
	running := true
	for running {
		utility.PrintAvailableCommands()
		command := utility.RequestUserCommand()
		command = strings.ToUpper(command)
		switch command {
		case "EXIT":
			running = false
		case "LIST":
			utility.ListAllPokemonNames(pokedex)
		case "ADD":
			pokemon := utility.RequestPokemon()
			utility.AddPokemon(pokemon.Name, pokemon, pokedex)
		case "GET":
			pokemonName := utility.RequestPokemonName()
			utility.GetPokemon(pokemonName, pokedex)
		case "BATTLE":
			pokemonOne, pokemonTwo := utility.GetPokemonForBattle(pokedex)
			service.Battle(pokemonOne, pokemonTwo)
		default:
			utility.PrintInvalidCommandMessage()
		}
	}
	utility.PrintExitMessage()
}
