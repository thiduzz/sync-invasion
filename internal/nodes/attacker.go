package nodes

import "github.com/thiduzz/code-kata-invasion/internal/constant"

type Attacker struct {
	Id       uint
	Name     string
	State    map[constant.AttackerState]bool
	Location *Location
}
