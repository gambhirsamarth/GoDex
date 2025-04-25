package utility

import (
	"GoDex/main/model"
	"strings"
)

func validateInputAndReturnPokemon(pokemonName string, pokedex map[string]model.Pokemon) (pokemon model.Pokemon) {
	if _, ok := pokedex[strings.ToUpper(pokemonName)]; ok {
		pokemon = pokedex[strings.ToUpper(pokemonName)]
	} else {
		PrintPokemonNotFoundMessage(pokemonName)
		validateInputAndReturnPokemon(RequestPokemonName(), pokedex)
	}
	return
}
