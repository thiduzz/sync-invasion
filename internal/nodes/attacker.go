package nodes

import "github.com/thiduzz/code-kata-invasion/internal/constant"

type Attacker struct {
	Id       uint
	Name     string
	State    map[constant.AttackerState]bool
	Location *Location
}

type AttackerInterface interface {
	NodeInterface
	IsDead() bool
	IsTrapped() bool
	Attack(*Location)
}

func NewAttacker(id uint, nameGeneratorFunc func() string) *Attacker {
	return &Attacker{Id: id, Name: nameGeneratorFunc()}
}

func (at *Attacker) IsDead() bool {
	return at.State[constant.Dead]
}

func (at *Attacker) IsTrapped() bool {
	return at.State[constant.Trapped]
}

func (at *Attacker) Attack(location *Location) {
	at.Location = location
}

func (at *Attacker) GetId() uint {
	return at.Id
}

func (at *Attacker) GetName() string {
	return at.Name
}
