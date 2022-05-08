package nodes

import "github.com/thiduzz/code-kata-invasion/tools"

type LocationCollection struct {
	Collection   map[uint]*Location
	ReferenceMap map[string]uint
}

func NewLocationCollection() *LocationCollection {
	return &LocationCollection{
		Collection:   make(map[uint]*Location),
		ReferenceMap: map[string]uint{},
	}
}

func (lc *LocationCollection) Add(location *Location) {
	lc.Collection[location.GetId()] = location
	lc.ReferenceMap[location.GetName()] = location.GetId()
}

func (lc *LocationCollection) GetById(id uint) *Location {
	return lc.Collection[id]
}

func (lc *LocationCollection) GetByName(name string) *Location {
	if referenceId, exists := lc.ReferenceMap[name]; exists {
		return lc.Collection[referenceId]
	}
	return nil
}

//GetUndestroyed O(N) - Return a slice of undestroyed city identifiers
func (lc *LocationCollection) GetUndestroyed() []uint {
	var locations []uint
	for _, location := range lc.Collection {
		if !location.IsDestroyed() {
			locations = append(locations, location.GetId())
		}
	}
	return locations
}

//GetRandom O(N) - Returns a random Location that is not destroyed
func (lc *LocationCollection) GetRandom(randomizer *tools.Randomizer) *Location {
	undestroyed := lc.GetUndestroyed()
	if len(undestroyed) <= 0 {
		return nil
	}
	randomizer.Shuffle(len(undestroyed), func(i, j int) {
		undestroyed[uint(i)], undestroyed[uint(i)] = undestroyed[uint(i)], undestroyed[uint(i)]
	})
	return lc.GetById(undestroyed[0])
}
