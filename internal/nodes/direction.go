package nodes

import "github.com/thiduzz/code-kata-invasion/internal/constant"

type Directions struct {
	Roads map[string][]uint
}

func NewDirectionCompass() map[string][]uint {
	return map[string][]uint{
		constant.DirectionNorth: nil,
		constant.DirectionSouth: nil,
		constant.DirectionWest:  nil,
		constant.DirectionEast:  nil,
	}
}

func (d *Directions) Add(direction string, location *Location, IsReverseDirection bool) {
	if IsReverseDirection {
		direction = d.InvertDirection(direction)
	}
	if !d.Exists(direction, location.GetId()) {
		d.Roads[direction] = append(d.Roads[direction], location.GetId())
	}
}

func (d *Directions) Exists(direction string, id uint) bool {
	for _, item := range d.Roads[direction] {
		if item == id {
			return true
		}
	}
	return false
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
