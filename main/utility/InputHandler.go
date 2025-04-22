package utility

import (
	"GoDex/main/model"
	"fmt"
)

func RequestUserCommand() (command string) {
	PrintRequestCommandMessage()
	fmt.Scanln(&command)
	return
}

func RequestPokemonName() (pokemonName string) {
	PrintRequestPokemonNameMessage()
	fmt.Scanln(&pokemonName)
	return
}

func RequestPokemonType() (pokemonType string) {
	PrintRequestPokemonTypeMessage()
	fmt.Scanln(&pokemonType)
	return
}

func RequestPokemonHP() (pokemonHP int) {
	PrintRequestPokemonHPMessage()
	fmt.Scanln(&pokemonHP)
	return
}

func RequestPokemonAttack() (pokemonAttack int) {
	PrintRequestPokemonAttackMessage()
	fmt.Scanln(&pokemonAttack)
	return
}

func RequestPokemonDefense() (pokemonDefense int) {
	PrintRequestPokemonDefenseMessage()
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
		Name:    name,
		Type:    pokeomonType,
		HP:      hp,
		Attack:  attack,
		Defense: defense,
	}
	return
}
