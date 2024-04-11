package players

import "fmt"

type warrior struct {
	weapon string
}

func (w warrior) Attack() {
	fmt.Println("Warrior is attacking with : ", w.weapon)
}

//Constructor
func NewWarrior(weapon string) IPlayer {
	return &warrior{weapon: weapon}
}
