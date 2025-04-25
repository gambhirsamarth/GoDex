package main

import (
	"GoDex/main/service/battle"
	"GoDex/main/service/commandexecutor"
	"GoDex/main/service/localstorage"
	"GoDex/main/utility/input"
	"GoDex/main/utility/output"
	"strings"
)

func main() {
	localstorage.InitializeDefaultPokedex()
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
			commandexecutor.ListAllPokemonNames()
		case "ADD":
			pokemon := input.RequestPokemon()
			localstorage.AddPokemon(pokemon)
		case "GET":
			pokemonName := strings.ToUpper(input.RequestPokemonName())
			commandexecutor.GetPokemon(pokemonName)
		case "BATTLE":
			pokemonOne, pokemonTwo := battle.GetPokemonForBattle()
			battle.Battle(pokemonOne, pokemonTwo)
		default:
			output.PrintInvalidCommandMessage()
		}
	}
	output.PrintExitMessage()
}
