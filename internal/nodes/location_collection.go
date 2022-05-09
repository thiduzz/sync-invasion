package nodes

import (
	"bytes"
	"fmt"
	"github.com/thiduzz/code-kata-invasion/tools"
	"sort"
)

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
	sort.Slice(locations, func(i, j int) bool { return locations[i] < locations[j] })
	return locations
}

//GetRandom O(N) - Returns a random Location that is not destroyed
func (lc *LocationCollection) GetRandom(randomizer *tools.Randomizer) *Location {
	undestroyed := lc.GetUndestroyed()
	if len(undestroyed) <= 0 {
		return nil
	}
	if randomizer != nil {
		randomizer.ShuffleUint(undestroyed)
	}
	return lc.GetById(undestroyed[0])
}

func (lc *LocationCollection) DestroyById(id uint) {
	for _, location := range lc.Collection {
		if location.GetId() == id {
			location.SetDestroyed(true)
		}
		location.DirectionsOutBound.Remove(id)
		location.DirectionsInBound.Remove(id)
	}
}

func (lc *LocationCollection) String() string {
	if len(lc.GetUndestroyed()) == 0 {
		return "The world has ended (all cities are destroyed)!"
	}
	var locationsBytes bytes.Buffer
	locationsBytes.WriteString(fmt.Sprintf("\n---Current World Map---\n\n"))
	notDestroyedLocations := lc.GetUndestroyed()
	for _, locationIdentifier := range notDestroyedLocations {
		if location := lc.GetById(locationIdentifier); location != nil {
			if directionString := location.DirectionsOutBound.GetDirectionString(lc, notDestroyedLocations); directionString != "" {
				locationsBytes.WriteString(fmt.Sprintf("%s%s\n", location.GetName(), directionString))
			} else {
				locationsBytes.WriteString(fmt.Sprintf("%s\n", location.GetName()))
			}
		}
	}
	return locationsBytes.String()
}
