package input

import (
	"GoDex/main/model"
	"GoDex/main/utility/output"
	"fmt"
	"strings"
)

func RequestUserCommand() (command string) {
	output.PrintRequestCommandMessage()
	fmt.Scanln(&command)
	return
}

func RequestPokemonName() (pokemonName string) {
	output.PrintRequestPokemonNameMessage()
	fmt.Scanln(&pokemonName)
	return
}

func RequestPokemonType() (pokemonType string) {
	output.PrintRequestPokemonTypeMessage()
	fmt.Scanln(&pokemonType)
	return
}

func RequestPokemonAttack() (pokemonAttack int) {
	output.PrintRequestPokemonAttackMessage()
	fmt.Scanln(&pokemonAttack)
	return
}

func RequestPokemonDefense() (pokemonDefense int) {
	output.PrintRequestPokemonDefenseMessage()
	fmt.Scanln(&pokemonDefense)
	return
}

func RequestPokemon() (pokemon model.Pokemon) {
	name := strings.ToUpper(RequestPokemonName())
	pokemonType := strings.ToUpper(RequestPokemonType())
	hp := 100
	attack := RequestPokemonAttack()
	defense := RequestPokemonDefense()
	pokemon = model.Pokemon{
		Name:    name,
		Type:    pokemonType,
		HP:      hp,
		Attack:  attack,
		Defense: defense,
	}
	return
}

