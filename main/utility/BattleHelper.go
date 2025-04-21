package utility

func CalculateDamage(attack, defense int) int {
	damage := attack - (defense / 2)
	if damage < 1 {
		damage = 1
	}
	return damage
}
