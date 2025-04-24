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
	pokemonType := RequestPokemonType()
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

func GetPokemonForBattle(pokedex map[string]model.Pokemon) (pokemonOne, pokemonTwo model.Pokemon) {
	pokemonOne = validateInputAndReturnPokemon(RequestPokemonName(), pokedex)
	pokemonTwo = validateInputAndReturnPokemon(RequestPokemonName(), pokedex)
	return
}
