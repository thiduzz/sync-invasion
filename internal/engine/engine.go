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
			fmt.Println("No more moves possible!")
			return nil
		}
	}
	fmt.Println("No more moves available!")
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
	var targetLocation, originalLocation *nodes.Location
	originalLocation = attacker.Location
	//check if it needs to initialize the attacker in a location
	if originalLocation == nil {
		//add entropy to the starting city definition
		en.Randomizer.Reseed()
		targetLocation = en.Locations.GetRandom(en.Randomizer)
		if targetLocation == nil {
			return nil, nil, errors.NewEngineErrorOp(errors.EndOfTheWorld)
		}
	} else {
		//acquire new target to attack
		targetLocation = en.planAttack(originalLocation)
		//Evaluate if it is trapped in the new city
		if targetLocation == nil {
			attacker.Trapped()
			return nil, nil, errors.NewEngineErrorOp(errors.AttackerTrapped)
		}
	}
	attacker.Attack(targetLocation)
	return targetLocation, originalLocation, nil
}

func (en *Engine) planAttack(originalLocation *nodes.Location) *nodes.Location {
	//Evaluate if it is trapped in the new city
	if originalLocation.DirectionsOutBound.IsDeadEnd() && originalLocation.DirectionsInBound.IsDeadEnd() {
		return nil
	}
	// Get all possible leaving the current location
	outboundRoadsAvailable := originalLocation.DirectionsOutBound.GetRandomizedRoads(en.Randomizer)
	for _, locationId := range outboundRoadsAvailable {
		//define as the destination the first random location that is not destroyed
		location := en.Locations.GetById(locationId)
		if !location.IsDestroyed() {
			return location
		}
	}
	// Get all possible leaving the current location
	inboundRoadsAvailable := originalLocation.DirectionsInBound.GetRandomizedRoads(en.Randomizer)
	for _, locationId := range inboundRoadsAvailable {
		//define as the destination the first random location that is not destroyed
		location := en.Locations.GetById(locationId)
		if !location.IsDestroyed() {
			return location
		}
	}
	return nil
}

//invade start attacking a location and possibly getting trapped or die in a fight
// Assumption: An attacker can only move through OutBound directions
func (en *Engine) invade(attacker *nodes.Attacker, target *nodes.Location) {
	if target != nil {
		//start invading new location
		if en.attacksInProgress[target.GetId()] == constant.EmptyCity {
			//Did not find anyone - come in
			fmt.Printf("Attacker %s has moved to %s\n", attacker, target.GetName())
			en.attacksInProgress[target.GetId()] = attacker.GetId()
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
