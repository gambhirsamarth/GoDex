package utility

import (
	"GoDex/main/model"
	"fmt"
)

func AddPokemon(name string, pokemon model.Pokemon, pokedex map[string]model.Pokemon) {
	if _, ok := pokedex[name]; ok {
		return
	} else {
		pokedex[name] = pokemon
	}
}

func GetPokemon(name string, pokedex map[string]model.Pokemon) {
	if _, ok := pokedex[name]; ok {
		fmt.Println(pokedex[name])
	} else {
		fmt.Println("Pokemon not found")
	}
}

func ListAllPokemonNames(pokedex map[string]model.Pokemon) {
	for _, pokemon := range pokedex {
		fmt.Println(pokemon.Name)
	}
}
