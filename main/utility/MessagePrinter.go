package utility

import (
	"fmt"
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
}

func PrintBattleWinner(winner, loser string) {
	fmt.Println()
	fmt.Printf("%s has fainted! %s wins!", loser, winner)
}
