package utility

import (
	"GoDex/main/model"
)

func CalculateDamage(attack, defense int) int {
	damage := attack - (defense / 2)
	if damage < 1 {
		damage = 1
	}
	return damage
}

func GetPokemonForBattle(pokedex map[string]model.Pokemon) (pokemonOne, pokemonTwo model.Pokemon) {
	pokemonOneName := RequestPokemonName()
	if _, okOne := pokedex[pokemonOneName]; okOne {
		pokemonOne = pokedex[pokemonOneName]
		pokemonTwoName := RequestPokemonName()
		if _, okTwo := pokedex[pokemonTwoName]; okTwo {
			pokemonTwo = pokedex[pokemonTwoName]
		} else {
			PrintPokemonNotFoundMessage(pokemonOneName)
			GetPokemonForBattle(pokedex)
		}
	} else {
		PrintPokemonNotFoundMessage(pokemonOneName)
		GetPokemonForBattle(pokedex)
	}
	return
}
