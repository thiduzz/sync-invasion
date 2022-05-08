package engine

import (
	"github.com/thiduzz/code-kata-invasion/internal/errors"
	"github.com/thiduzz/code-kata-invasion/internal/nodes"
	"github.com/thiduzz/code-kata-invasion/tools"
)

type Engine struct {
	Locations      *nodes.LocationCollection
	Attackers      *nodes.AttackerCollection
	Randomizer     *tools.Randomizer
	AttackersQty   uint
	MaxMoves       uint
	onGoingAttacks map[uint]uint
}

func NewEngine(locations *nodes.LocationCollection, randomizer *tools.Randomizer, attackersQty uint, maxMoves uint) *Engine {
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
		en.Randomizer.Reseed()
		orderOfAttack := en.Attackers.Sort(en.Randomizer)
		for _, attackerIdentifier := range orderOfAttack {
			attacker := en.Attackers.GetById(attackerIdentifier)
			//method responsible for acquiring the target and evaluating attackers state ("health")
			target, err := en.attack(attacker)
			if err != nil {
				return err
			}
			//method responsible for occupying the location
			en.invade(attacker, target)
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

//attack define the target that the current attacker will invade
func (en *Engine) attack(attacker *nodes.Attacker) (*nodes.Location, *nodes.Location, error) {
	//abort if the attacker is dead (no-action)
	if attacker.IsDead() {
		return nil, nil, errors.NewEngineErrorOp(errors.AttackerDead)
	}
	//abort if the attacker is trapped (no-action)
	if attacker.IsTrapped() {
		return nil, nil, errors.NewEngineErrorOp(errors.AttackerTrapped)
	}
	var newLocation, originalLocation *nodes.Location
	originalLocation = attacker.Location
	//check if it needs to initialize the attacker in a location
	if originalLocation == nil {
		//add entropy to the starting city definition
		en.Randomizer.Reseed()
		newLocation = en.Locations.GetRandom(en.Randomizer)
		if newLocation == nil {
			return nil, nil, errors.NewEngineErrorOp(errors.EndOfTheWorld)
		}
	} else {
		//TODO: pick a city to attack from the outbound routes
	}
	attacker.Attack(newLocation)
	return newLocation, originalLocation, nil
}

//invade start attacking a location by
func (en *Engine) invade(attacker *nodes.Attacker, target *nodes.Location) {
	if target != nil {
		//start invading new location
		if en.attacksInProgress[target.GetId()] == constant.EmptyCity {
			//Did not find anyone - come in
		}
	}

func (en *Engine) invade(attacker *nodes.Attacker, target *nodes.Location) {

}
