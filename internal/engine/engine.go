package engine

import (
	"github.com/thiduzz/code-kata-invasion/internal/errors"
	"github.com/thiduzz/code-kata-invasion/internal/nodes"
	"math/rand"
)

type Engine struct {
	Locations         *nodes.LocationCollection
	Attackers         *nodes.AttackerCollection
	EntropyRandomizer *rand.Rand
	AttackersQty      uint
	MaxMoves          uint
}

func NewEngine(locations *nodes.LocationCollection, randomizer *rand.Rand, attackersQty uint, maxMoves uint) *Engine {
	return &Engine{
		Locations:         locations,
		AttackersQty:      attackersQty,
		MaxMoves:          maxMoves,
		Attackers:         nodes.NewAttackerCollection(),
		EntropyRandomizer: randomizer,
	}
}

func (en *Engine) Start() error {
	return nil
}

//PrepareAttackers Generate aliens with a factory and add them to the engine to be later on "worked"
func (en *Engine) PrepareAttackers(factory nodes.AttackerFactoryInterface) error {
	for i := uint(1); i <= en.AttackersQty; i++ {
		attacker, err := factory.Generate(nodes.Attacker{}, i)
		if err != nil {
			return errors.NewEngineErrorWrap("alienFactory", err)
		}
		en.Attackers.Add(attacker)
	}
	return nil
}
