package service

import "GoDex/main/model"

var pokedex = map[string]model.Pokemon{}

func GetPokedex() map[string]model.Pokemon {
	return pokedex
}

func InitializeDefaultPokedex() {
	pokedex = map[string]model.Pokemon{
		"PIKACHU": {
			Name:    "PIKACHU",
			Type:    "ELECTRIC",
			HP:      100,
			Attack:  55,
			Defense: 40,
		},
		"CHARMANDER": {
			Name:    "CHARMANDER",
			Type:    "FIRE",
			HP:      100,
			Attack:  52,
			Defense: 43,
		},
		"BULBASAUR": {
			Name:    "BULBASAUR",
			Type:    "GRASS",
			HP:      100,
			Attack:  49,
			Defense: 49,
		},
		"SQUIRTLE": {
			Name:    "SQUIRTLE",
			Type:    "WATER",
			HP:      100,
			Attack:  48,
			Defense: 65,
		},
		"EEVEE": {
			Name:    "EEVEE",
			Type:    "NORMAL",
			HP:      100,
			Attack:  55,
			Defense: 50,
		},
		"JIGGLYPUFF": {
			Name:    "JIGGLYPUFF",
			Type:    "FAIRY",
			HP:      100,
			Attack:  45,
			Defense: 20,
		},
		"MEOWTH": {
			Name:    "MEOWTH",
			Type:    "NORMAL",
			HP:      100,
			Attack:  45,
			Defense: 35,
		},
	}
}
