package engine

import (
	"github.com/thiduzz/code-kata-invasion/internal/errors"
	"github.com/thiduzz/code-kata-invasion/internal/nodes"
	"math/rand"
	"time"
)

type Engine struct {
	Locations    *nodes.LocationCollection
	Attackers    *nodes.AttackerCollection
	Randomizer   *rand.Rand
	AttackersQty uint
	MaxMoves     uint
}

func NewEngine(locations *nodes.LocationCollection, randomizer *rand.Rand, attackersQty uint, maxMoves uint) *Engine {
	return &Engine{
		Locations:    locations,
		AttackersQty: attackersQty,
		MaxMoves:     maxMoves,
		Attackers:    nodes.NewAttackerCollection(),
		Randomizer:   randomizer,
	}
}

//Start This function defines the loop which will be running workers and locations
// according to the specification - which is "👾👾 Alien Invasion 👾👾"
func (en *Engine) Start() error {
	for iterations := uint(0); iterations < en.MaxMoves; iterations++ {
		//add entropy to the attacker order
		en.reseedRandom()
		orderOfAttack := en.Attackers.Sort(en.Randomizer)
		for _, attackerIdentifier := range orderOfAttack {
			attacker := en.Attackers.GetById(attackerIdentifier)
			err := en.attack(attacker)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (en *Engine) attack(attacker *nodes.Attacker) error {
	//abort if the attacker is dead (no-action)
	if attacker.IsDead() {
		return errors.NewEngineErrorOp(errors.AttackerDead)
	}
	//abort if the attacker is trapped (no-action)
	if attacker.IsTrapped() {
		return errors.NewEngineErrorOp(errors.AttackerTrapped)
	}

	//check if it needs to initialize the attacker in a location
	if attacker.Location == nil {
		//add entropy to the starting city definition
		en.reseedRandom()
		if location := en.Locations.GetRandom(en.Randomizer); location == nil {
			return errors.NewEngineErrorOp(errors.EndOfTheWorld)
		}
	}
	return nil
}

//PrepareAttackers Generate aliens with a factory and add them to the engine to be later on "worked"
func (en *Engine) PrepareAttackers(factory nodes.AttackerFactoryInterface) error {
	for i := uint(1); i <= en.AttackersQty; i++ {
		attacker, err := factory.Generate(nodes.Attacker{}, i)
		if err != nil {
			return errors.NewEngineErrorWrap(errors.AttackerFactory, err)
		}
		en.Attackers.Add(attacker)
	}
	return nil
}

//reseedRandom Reseed the randomizer with the current time providing a different entropy
func (en *Engine) reseedRandom() {
	en.Randomizer.Seed(time.Now().UnixMilli())
}
