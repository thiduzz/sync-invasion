package nodes

import "github.com/thiduzz/code-kata-invasion/tools"

type AttackerCollection struct {
	Collection   map[uint]*Attacker
	ReferenceMap map[string]uint
	keys         []uint
}

func NewAttackerCollection() *AttackerCollection {
	return &AttackerCollection{
		Collection:   make(map[uint]*Attacker),
		ReferenceMap: map[string]uint{},
	}
}

//Add Adds attacker to the list of tracked attackers
func (ac *AttackerCollection) Add(attacker *Attacker) {
	//add to the main map (which holds the reference to the attacker)
	ac.Collection[attacker.GetId()] = attacker
	//add to the support map that allows query of locations per name
	ac.ReferenceMap[attacker.GetName()] = attacker.GetId()
	//add to slice that enable a quick access of all ids removing the need to iterate over a map
	ac.keys = append(ac.keys, attacker.GetId())
}

//GetById Get an attacker by its id
func (ac *AttackerCollection) GetById(id uint) *Attacker {
	return ac.Collection[id]
}

//GetByName Gets an attacker by its name using a support link map
func (ac *AttackerCollection) GetByName(name string) *Attacker {
	if referenceId, exists := ac.ReferenceMap[name]; exists {
		return ac.Collection[referenceId]
	}
	return nil
}

//Sort Given that maps are notoriously "semi-random" this functions
// ensure to add the necessary entropy so that the attackers will act in
// different order in most iterations - it accepts a randomizer that
// enable a better control over the undeterministic behavior while testing
func (ac *AttackerCollection) Sort(randomizer *tools.Randomizer) []uint {
	randomizer.Shuffle(len(ac.keys), func(i, j int) {
		ac.keys[i], ac.keys[j] = ac.keys[j], ac.keys[i]
	})
	return ac.keys
}
