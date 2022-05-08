package engine

import (
	"github.com/thiduzz/code-kata-invasion/internal/errors"
	"github.com/thiduzz/code-kata-invasion/internal/nodes"
)

type Engine struct {
	Locations    *nodes.LocationCollection
	Attackers    *nodes.AttackerCollection
	AttackersQty uint
	MaxMoves     uint
}

func NewEngine(locations *nodes.LocationCollection, attackersQty uint, maxMoves uint) *Engine {
	return &Engine{Locations: locations, AttackersQty: attackersQty, MaxMoves: maxMoves}
}

func (en *Engine) Start() error {
	return nil
}

func (en *Engine) PrepareAttackers(factory *nodes.AttackerFactory) error {
	for i := uint(1); i <= en.AttackersQty; i++ {
		attacker, err := factory.Attacker(nodes.Attacker{}, i)
		if err != nil {
			return errors.NewEngineErrorWrap("alienFactory", err)
		}
	}
	return nil
}
