package utility

import (
	"GoDex/main/model"
	"fmt"
)

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
	return
}

func RequestPokemonType() (pokemonType string) {
	fmt.Println()
	fmt.Println("Please enter the Pokemon type: ")
	fmt.Scanln(&pokemonType)
	return
}

func RequestPokemonHP() (pokemonHP int) {
	fmt.Println()
	fmt.Println("Please enter the Pokemon HP: ")
	fmt.Scanln(&pokemonHP)
	return
}

func RequestPokemonAttack() (pokemonAttack int) {
	fmt.Println()
	fmt.Println("Please enter the Pokemon Attack: ")
	fmt.Scanln(&pokemonAttack)
	return
}

func RequestPokemonDefense() (pokemonDefense int) {
	fmt.Println()
	fmt.Println("Please enter the Pokemon Defense: ")
	fmt.Scanln(&pokemonDefense)
	return
}

func RequestPokemon() (pokemon model.Pokemon) {
	name := RequestPokemonName()
	pokeomonType := RequestPokemonType()
	hp := RequestPokemonHP()
	attack := RequestPokemonAttack()
	defense := RequestPokemonDefense()
	pokemon = model.Pokemon{
		Name: name,
		Type: pokeomonType,
		HP: hp,
		Attack: attack,
		Defense: defense,
	}
	return
}
