package commandexecutor

import (
	"GoDex/main/service/localstorage"
	"GoDex/main/utility/output"
)

func GetPokemon(name string) {
	if _, ok := localstorage.Pokedex[name]; ok {
		output.PrintPokemonModel(localstorage.Pokedex[name])
	} else {
		output.PrintPokemonNotFoundMessage(name)
	}
}

func ListAllPokemonNames() {
	for _, pokemon := range localstorage.Pokedex {
		output.PrintPokemonName(pokemon.Name)
	}
}
