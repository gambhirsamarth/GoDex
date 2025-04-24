package utility

import (
	"GoDex/main/model"
	"fmt"
	"time"
)

func PrintWelcomeMessage() {
	fmt.Println()
	fmt.Println("Welcome to GoDex!")
}

func PrintAvailableCommands() {
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println()
	fmt.Println("ADD : to add a Pokemon")
	fmt.Println("LIST : to list all Pokemon names")
	fmt.Println("GET : to get pokemon details")
	fmt.Println("BATTLE : to start a battle between two Pokemons")
	fmt.Println("EXIT : to exit GoDex")
}

func PrintInvalidCommandMessage() {
	fmt.Println()
	fmt.Println("Invalid command")
}

func PrintExitMessage() {
	fmt.Println()
	fmt.Println("Exiting GoDex")
}

func PrintBattleStartMessage(pokemonOne, pokemonTwo string) {
	fmt.Println()
	fmt.Printf("Battle Start: %s vs %s\n", pokemonOne, pokemonTwo)
}

func PrintAttackMessage(attackerName, defenderName string, damage, hp int) {
	fmt.Println()
	fmt.Printf("%s attacks %s for %d damage. %s's HP is now %d.", attackerName, defenderName, damage, defenderName, hp)
	time.Sleep(time.Second)
}

func PrintBattleWinner(loser, winner string) {
	fmt.Println()
	fmt.Printf("%s has fainted! %s wins!", loser, winner)
}

func PrintRequestCommandMessage() {
	fmt.Println()
	fmt.Println("Please enter a command: ")
}

func PrintRequestPokemonNameMessage() {
	fmt.Println()
	fmt.Println("Please enter the Pokemon name: ")
}

func PrintRequestPokemonTypeMessage() {
	fmt.Println()
	fmt.Println("Please enter the Pokemon type: ")
}

func PrintRequestPokemonHPMessage() {
	fmt.Println()
	fmt.Println("Please enter the Pokemon HP: ")
}

func PrintRequestPokemonAttackMessage() {
	fmt.Println()
	fmt.Println("Please enter the Pokemon attack: ")
}

func PrintRequestPokemonDefenseMessage() {
	fmt.Println()
	fmt.Println("Please enter the Pokemon defense: ")
}

func PrintPokemonNotFoundMessage(pokemonName string) {
	fmt.Println()
	fmt.Printf("Pokemon %s not found", pokemonName)
}

func PrintPokemonModel(pokemon model.Pokemon) {
	fmt.Println(pokemon)
}

func printPokemonName(pokemonName string) {
	fmt.Println()
	fmt.Printf("Pokemon name: %s", pokemonName)
}
