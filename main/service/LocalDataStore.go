package service

import "GoDex/main/model"

var pokedex = map[string]model.Pokemon{}

func GetPokedex() map[string]model.Pokemon {
	return pokedex
}

func InitializeDefaultPokedex() {
	pokedex = map[string]model.Pokemon{
		"Pikachu": {
			Name:    "Pikachu",
			Type:    "Electric",
			HP:      100,
			Attack:  55,
			Defense: 40,
		},
		"Charmander": {
			Name:    "Charmander",
			Type:    "Fire",
			HP:      100,
			Attack:  52,
			Defense: 43,
		},
		"Bulbasaur": {
			Name:    "Bulbasaur",
			Type:    "Grass",
			HP:      100,
			Attack:  49,
			Defense: 49,
		},
		"Squirtle": {
			Name:    "Squirtle",
			Type:    "Water",
			HP:      100,
			Attack:  48,
			Defense: 65,
		},
		"Eevee": {
			Name:    "Eevee",
			Type:    "Normal",
			HP:      100,
			Attack:  55,
			Defense: 50,
		},
		"Jigglypuff": {
			Name:    "Jigglypuff",
			Type:    "Fairy",
			HP:      100,
			Attack:  45,
			Defense: 20,
		},
		"Meowth": {
			Name:    "Meowth",
			Type:    "Normal",
			HP:      100,
			Attack:  45,
			Defense: 35,
		},
	}
}
