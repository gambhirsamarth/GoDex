package utility

import "GoDex/main/model"

func validateInputAndReturnPokemon(pokemonName string, pokedex map[string]model.Pokemon) (pokemon model.Pokemon) {
	if _, ok := pokedex[pokemonName]; ok {
		pokemon = pokedex[pokemonName]
	} else {
		PrintPokemonNotFoundMessage(pokemonName)
		validateInputAndReturnPokemon(RequestPokemonName(), pokedex)
	}
	return
}
