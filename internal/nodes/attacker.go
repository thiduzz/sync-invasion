package nodes

import (
	"fmt"
	"github.com/thiduzz/code-kata-invasion/internal/constant"
)

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
	Die()
	Trapped()
}

func NewAttacker(id uint, nameGeneratorFunc func() string) *Attacker {
	return &Attacker{Id: id, Name: nameGeneratorFunc(), State: map[constant.AttackerState]bool{}}
}

func (at *Attacker) Die() {
	at.State[constant.Dead] = true
}

func (at *Attacker) Trapped() {
	at.State[constant.Trapped] = true
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

func (at *Attacker) IsDead() bool {
	return at.State[constant.Dead]
}

func (at *Attacker) IsTrapped() bool {
	return at.State[constant.Trapped]
}

func (at *Attacker) String() string {
	return fmt.Sprintf("#%d (%s)", at.GetId(), at.GetName())
}
