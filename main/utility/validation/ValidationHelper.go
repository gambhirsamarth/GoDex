package validation

import (
	"GoDex/main/model"
	"GoDex/main/utility/input"
	"GoDex/main/utility/output"
	"strings"
)

func ValidateInputAndReturnPokemon(pokemonName string, pokedex map[string]model.Pokemon) (pokemon model.Pokemon) {
	if CheckPokemonExistsInPokedex(pokemonName, pokedex) {
		pokemon = pokedex[strings.ToUpper(pokemonName)]
	} else {
		output.PrintPokemonNotFoundMessage(pokemonName)
		ValidateInputAndReturnPokemon(input.RequestPokemonName(), pokedex)
	}
	return
}

func CheckPokemonExistsInPokedex(name string, pokedex map[string]model.Pokemon) bool {
	if _, ok := pokedex[strings.ToUpper(name)]; ok {
		return true
	} else {
		return false
	}
}
