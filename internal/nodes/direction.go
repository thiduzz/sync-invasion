package nodes

import (
	"bytes"
	"fmt"
	"github.com/thiduzz/code-kata-invasion/internal/constant"
	"github.com/thiduzz/code-kata-invasion/tools"
)

type Directions struct {
	Roads map[string]map[uint]bool
}

//NewDirectionCompass Provides a default compass map
func NewDirectionCompass() map[string]map[uint]bool {
	return map[string]map[uint]bool{
		constant.DirectionNorth: {},
		constant.DirectionSouth: {},
		constant.DirectionWest:  {},
		constant.DirectionEast:  {},
	}
}

//Add creates a link between cities - so that the attacker can traverse it
func (d *Directions) Add(direction string, location *Location, IsReverseDirection bool) {
	// defines if it should reverse the direction - useful when defining inbound routes
	if IsReverseDirection {
		direction = d.InvertDirection(direction)
	}
	// Adds reference of location to Roads Direction (only possible value is "true" or it doesnt exist)
	if !d.Exists(direction, location.GetId()) {
		d.Roads[direction][location.GetId()] = true
	}
}

func (d *Directions) Exists(direction string, id uint) bool {
	return d.Roads[direction][id]
}

// GetRandomizedRoads GetRandom O(N) - Returns random location Ids attached to this direction
func (d *Directions) GetRandomizedRoads(randomizer *tools.Randomizer) []uint {
	var locationIds []uint
	for dir, _ := range d.Roads {
		// if it doesnt have roads in this direction skip iteration
		if len(d.Roads[dir]) > 0 {
			continue
		}
		// adds reference to locations in a slice (easier to handle and shuffle)
		for u, _ := range d.Roads[dir] {
			locationIds = append(locationIds, u)
		}
	}
	if len(locationIds) <= 0 {
		return locationIds
	}
	randomizer.ShuffleUint(locationIds)
	return locationIds
}

//GetDirectionString Returns a string containing all the directions and references in the expected stdout format
func (d *Directions) GetDirectionString(collection *LocationCollection, undestroyedIds []uint) string {
	var directionsBytes bytes.Buffer
	for direction, locationIdsMap := range d.Roads {
		if len(locationIdsMap) <= 0 {
			continue
		}
		for _, locationIdentifier := range undestroyedIds {
			if locationIdsMap[locationIdentifier] {
				if location := collection.GetById(locationIdentifier); location != nil {
					directionsBytes.WriteString(fmt.Sprintf(" %s=%s", direction, location.GetName()))
				}
			}
		}
	}
	return directionsBytes.String()
}

//InvertDirection Method responsible for reversing the compass - utilized when adding InBound connection
//The north of the OutBound location is the south of the InBound location
func (d *Directions) InvertDirection(direction string) string {
	switch direction {
	case constant.DirectionNorth:
		return constant.DirectionSouth
	case constant.DirectionSouth:
		return constant.DirectionNorth
	case constant.DirectionWest:
		return constant.DirectionEast
	case constant.DirectionEast:
		return constant.DirectionWest
	}
	return ""
}

//IsDeadEnd checks if there are any leading road available
func (d *Directions) IsDeadEnd() bool {
	for dir, _ := range d.Roads {
		if len(d.Roads[dir]) > 0 {
			return false
		}
	}
	return true
}

//Remove Unset all directions that point to a destroyed location
func (d *Directions) Remove(id uint) {
	for dir, _ := range d.Roads {
		if _, exists := d.Roads[dir][id]; exists {
			delete(d.Roads[dir], id)
		}
	}
}
