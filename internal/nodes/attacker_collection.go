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

func (ac *AttackerCollection) Add(attacker *Attacker) {
	ac.Collection[attacker.GetId()] = attacker
	ac.ReferenceMap[attacker.GetName()] = attacker.GetId()
	ac.keys = append(ac.keys, attacker.GetId())
}

func (ac *AttackerCollection) GetById(id uint) *Attacker {
	return ac.Collection[id]
}

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
