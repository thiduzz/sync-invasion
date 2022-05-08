package nodes

import "github.com/thiduzz/code-kata-invasion/internal/constant"

type Attacker struct {
	Id       uint
	Name     string
	State    map[constant.AttackerState]bool
	Location *Location
}

func NewAttacker(id uint, nameGeneratorFunc func() string) *Attacker {
	return &Attacker{Id: id, Name: nameGeneratorFunc()}
}
