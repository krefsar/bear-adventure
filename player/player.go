package player

import (
	"github.com/krefsar/bear-adventure/powers"
)

type Player struct {
	Name string
	Powers []*powers.Power
}

func (p *Player) GainPower(power *powers.Power) {
	p.Powers = append(p.Powers, power)
}
