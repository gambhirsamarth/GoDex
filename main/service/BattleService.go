package service

import (
	"GoDex/main/model"
	"GoDex/main/utility/input"
	"GoDex/main/utility/output"
	"GoDex/main/utility/validation"
)

func Battle(p1, p2 model.Pokemon) {
	output.PrintBattleStartMessage(p1.Name, p2.Name)

	for p1.HP > 0 && p2.HP > 0 {
		// p1 attacks p2
		damageToP2 := CalculateDamage(p1.Attack, p2.Defense)
		p2.HP -= damageToP2
		if p2.HP < 0 {
			p2.HP = 0
		}
		output.PrintAttackMessage(p1.Name, p2.Name, damageToP2, p2.HP)

		if p2.HP <= 0 {
			output.PrintBattleWinner(p2.Name, p1.Name)
			break
		}

		// p2 attacks p1
		damageToP1 := CalculateDamage(p2.Attack, p1.Defense)
		p1.HP -= damageToP1
		if p1.HP < 0 {
			p1.HP = 0
		}
		output.PrintAttackMessage(p2.Name, p1.Name, damageToP1, p1.HP)

		if p1.HP <= 0 {
			output.PrintBattleWinner(p1.Name, p2.Name)
			break
		}
	}
}

func CalculateDamage(attack, defense int) int {
	damage := attack - (defense / 2)
	if damage < 1 {
		damage = 1
	}
	return damage
}

func GetPokemonForBattle() (pokemonOne, pokemonTwo model.Pokemon) {
	pokemonOne = validation.ValidateInputAndReturnPokemon(input.RequestPokemonName(), Pokedex)
	pokemonTwo = validation.ValidateInputAndReturnPokemon(input.RequestPokemonName(), Pokedex)
	return
}
