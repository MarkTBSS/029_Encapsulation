package main

import "github.com/MarkTBSS/029_Encapsulation/players"

func playerAttack(player players.IPlayer) {
	player.Attack()
}

func main() {
	aragon := players.NewWarrior("Knight Sword")
	playerAttack(aragon)

	gandalf := players.NewMage("Fire Ball")
	playerAttack(gandalf)
}
