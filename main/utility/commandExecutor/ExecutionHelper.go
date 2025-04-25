package commandExecutor

import (
	"GoDex/main/service"
	"GoDex/main/utility/output"
)

func GetPokemon(name string) {
	if _, ok := service.Pokedex[name]; ok {
		output.PrintPokemonModel(service.Pokedex[name])
	} else {
		output.PrintPokemonNotFoundMessage(name)
	}
}

func ListAllPokemonNames() {
	for _, pokemon := range service.Pokedex {
		output.PrintPokemonName(pokemon.Name)
	}
}
