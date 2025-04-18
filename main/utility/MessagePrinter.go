package utility

import "fmt"

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
