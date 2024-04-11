package players

import "fmt"

type mage struct {
	spell string
}

func (m mage) Attack() {
	fmt.Println("Warrior is attacking with : ", m.spell)
}

func NewMage(spell string) IPlayer {
	return &mage{spell: spell}
}
