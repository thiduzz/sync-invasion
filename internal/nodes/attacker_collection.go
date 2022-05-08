package nodes

type AttackerCollection struct {
	Collection   map[uint]*Attacker
	ReferenceMap map[string]uint
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
