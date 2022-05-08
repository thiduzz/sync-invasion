package nodes

import "github.com/thiduzz/code-kata-invasion/internal/constant"

type Location struct {
	Id                 uint
	Name               string
	State              map[constant.LocationState]bool
	DirectionsOutBound Directions
	DirectionsInBound  Directions
}

func NewLocation(id uint, name string) *Location {
	return &Location{
		Id:   id,
		Name: name,
		DirectionsOutBound: Directions{
			Roads: NewDirectionCompass(),
		},
		DirectionsInBound: Directions{
			Roads: NewDirectionCompass(),
		},
	}
}

type LocationInterface interface {
	NodeInterface
	IsDestroyed() bool
}

func (l *Location) GetId() uint {
	return l.Id
}

func (l *Location) GetName() string {
	return l.Name
}

func (l *Location) IsDestroyed() bool {
	return l.State[constant.Destroyed]
}
