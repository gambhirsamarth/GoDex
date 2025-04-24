package utility

import (
	"GoDex/main/model"
)

func AddPokemon(name string, pokemon model.Pokemon, pokedex map[string]model.Pokemon) {
	ValidateAndAddPokemon(name, pokemon, pokedex)
}

func GetPokemon(name string, pokedex map[string]model.Pokemon) {
	if _, ok := pokedex[name]; ok {
		PrintPokemonModel(pokedex[name])
	} else {
		PrintPokemonNotFoundMessage(name)
	}
}

func ListAllPokemonNames(pokedex map[string]model.Pokemon) {
	for _, pokemon := range pokedex {
		printPokemonName(pokemon.Name)
	}
}

func ValidateAndAddPokemon(name string, pokemon model.Pokemon, pokedex map[string]model.Pokemon) {
	if _, ok := pokedex[name]; ok {
		return
	} else {
		pokedex[name] = pokemon
	}
}
