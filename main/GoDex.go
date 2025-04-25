package main

import (
	"GoDex/main/service"
	"GoDex/main/utility/commandExecutor"
	"GoDex/main/utility/input"
	"GoDex/main/utility/output"
	"strings"
)

func main() {
	service.InitializeDefaultPokedex()
	output.PrintWelcomeMessage()
	running := true
	for running {
		output.PrintAvailableCommands()
		command := input.RequestUserCommand()
		command = strings.ToUpper(command)
		switch command {
		case "EXIT":
			running = false
		case "LIST":
			commandExecutor.ListAllPokemonNames()
		case "ADD":
			pokemon := input.RequestPokemon()
			service.AddPokemon(pokemon)
		case "GET":
			pokemonName := strings.ToUpper(input.RequestPokemonName())
			commandExecutor.GetPokemon(pokemonName)
		case "BATTLE":
			pokemonOne, pokemonTwo := service.GetPokemonForBattle()
			service.Battle(pokemonOne, pokemonTwo)
		default:
			output.PrintInvalidCommandMessage()
		}
	}
	output.PrintExitMessage()
}
