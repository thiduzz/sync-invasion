package engine

import (
	goerrors "errors"
	"fmt"
	"github.com/thiduzz/code-kata-invasion/internal/constant"
	"github.com/thiduzz/code-kata-invasion/internal/errors"
	"github.com/thiduzz/code-kata-invasion/internal/nodes"
	"github.com/thiduzz/code-kata-invasion/tools"
)

type Engine struct {
	Locations         *nodes.LocationCollection
	Attackers         *nodes.AttackerCollection
	Randomizer        *tools.Randomizer
	AttackersQty      uint
	MaxMoves          uint
	attacksInProgress map[uint]uint
}

func NewEngine(locations *nodes.LocationCollection, randomizer *tools.Randomizer, attackersQty uint, maxMoves uint) *Engine {
	return &Engine{
		Locations:         locations,
		AttackersQty:      attackersQty,
		MaxMoves:          maxMoves,
		Attackers:         nodes.NewAttackerCollection(),
		Randomizer:        randomizer,
		attacksInProgress: map[uint]uint{},
	}
}

//Start This function defines the loop which will be running workers and locations
// according to the specification - which is "👾👾 Alien Invasion 👾👾"
func (en *Engine) Start() error {
	for iterations := uint(0); iterations < en.MaxMoves; iterations++ {
		//add entropy to the attacker order
		en.Randomizer.Reseed()
		orderOfAttack := en.Attackers.Sort(en.Randomizer)
		attacksHappened := false
		for _, attackerIdentifier := range orderOfAttack {
			attacker := en.Attackers.GetById(attackerIdentifier)
			//method responsible for acquiring the target and evaluating attackers state ("health")
			target, originalLocation, err := en.attack(attacker)
			if err != nil {
				if shouldInterrupt, noOperations := evaluateError(err); noOperations {
					continue
				} else if shouldInterrupt {
					return err
				}
			}
			attacksHappened = true
			//stop invading current location
			if originalLocation != nil && en.attacksInProgress[originalLocation.GetId()] == attacker.GetId() {
				delete(en.attacksInProgress, originalLocation.GetId())
			}
			//start invading new location
			en.invade(attacker, target)
		}
		// abort engine
		if !attacksHappened {
			return nil
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

//invade start attacking a location and possibly getting trapped or die in a fight
// Assumption: An attacker can only move through OutBound directions
func (en *Engine) invade(attacker *nodes.Attacker, target *nodes.Location) {
	if target != nil {
		//start invading new location
		if en.attacksInProgress[target.GetId()] == constant.EmptyCity {
			//Did not find anyone - come in
			en.attacksInProgress[target.GetId()] = attacker.GetId()
			//Evaluate if it is trapped in the new city
			if target.DirectionsOutBound.IsDeadEnd() {
				attacker.Trapped()
			}
		} else {
			//Found another attacker - Fight!
			enemy := en.Attackers.GetById(en.attacksInProgress[target.GetId()])
			en.Locations.DestroyById(target.GetId())
			enemy.Die()
			attacker.Die()
			fmt.Printf("%s has been destroyed by alien %s and %s\n", target.GetName(), enemy, attacker)
		}
	}
}

//evaluateError Determine how a specific error is treated on the lifecycle of the engine
func evaluateError(err error) (interrupt bool, noOperation bool) {
	var ee *errors.EngineError
	if goerrors.As(err, &ee) {
		switch ee.Op {
		case errors.AttackerDead, errors.AttackerTrapped:
			return false, true
		}
	}
	return true, false
}
