package service

import (
	"GoDex/main/model"
	"GoDex/main/utility"
	
)

func Battle(p1, p2 model.Pokemon) {
	fmt.Printf("Battle Start: %s vs %s\n", p1.Name, p2.Name)

	for p1.HP > 0 && p2.HP > 0 {
		// p1 attacks p2
		damageToP2 := utility.CalculateDamage(p1.Attack, p2.Defense)
		p2.HP -= damageToP2
		utility.PrintAttackMessage(p1.Name, p2.Name, damageToP2, p2.HP)

		if p2.HP <= 0 {
			utility.PrintBattleWinner(p2.Name, p1.Name)
			break
		}

		// p2 attacks p1
		damageToP1 := utility.CalculateDamage(p2.Attack, p1.Defense)
		p1.HP -= damageToP1
		utility.PrintAttackMessage(p2.Name, p1.Name, damageToP1, p1.HP)

		if p1.HP <= 0 {
			utility.PrintBattleWinner(p1.Name, p2.Name)
			break
		}
	}
}
