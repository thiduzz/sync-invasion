package engine

import "github.com/thiduzz/code-kata-invasion/internal/nodes"

type Engine struct {
	Locations    *nodes.LocationCollection
	AttackersQty uint
	MaxMoves     uint
}

func NewEngine(locations *nodes.LocationCollection, attackersQty uint, maxMoves uint) *Engine {
	return &Engine{Locations: locations, AttackersQty: attackersQty, MaxMoves: maxMoves}
}

func (en *Engine) Start() error {
	return nil
}

func (en *Engine) PrepareAttackers() {

}
