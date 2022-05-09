package nodes

import (
	"bytes"
	"fmt"
	"github.com/thiduzz/code-kata-invasion/internal/constant"
)

type Directions struct {
	Roads map[string]map[uint]bool
}

func NewDirectionCompass() map[string]map[uint]bool {
	return map[string]map[uint]bool{
		constant.DirectionNorth: {},
		constant.DirectionSouth: {},
		constant.DirectionWest:  {},
		constant.DirectionEast:  {},
	}
}

func (d *Directions) Add(direction string, location *Location, IsReverseDirection bool) {
	if IsReverseDirection {
		direction = d.InvertDirection(direction)
	}
	if !d.Exists(direction, location.GetId()) {
		d.Roads[direction][location.GetId()] = true
	}
}

func (d *Directions) Exists(direction string, id uint) bool {
	return d.Roads[direction][id]
}

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
