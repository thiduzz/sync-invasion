package nodes

import "github.com/thiduzz/code-kata-invasion/internal/constant"

type Location struct {
	Id         uint
	Name       string
	State      map[constant.LocationState]bool
	Directions Directions
}

type Directions struct {
	Roads     map[string]*Location
	Blueprint []string
}

func NewLocation(id uint, name string, originalDirectionsBlueprint []string) Location {
	return Location{
		Id:   id,
		Name: name,
		Directions: Directions{
			Roads: map[string]*Location{
				constant.DirectionNorth: nil,
				constant.DirectionSouth: nil,
				constant.DirectionWest:  nil,
				constant.DirectionEast:  nil,
			},
			Blueprint: originalDirectionsBlueprint,
		}}
}

type LocationInterface interface {
	GetId() uint
	GetName() string
}

func (l Location) GetId() uint {
	return l.Id
}

func (l Location) GetName() string {
	return l.Name
}
