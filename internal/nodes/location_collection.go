package nodes

import "log"

type LocationCollection struct {
	Collection map[uint]*Location
}

func NewLocationCollection() *LocationCollection {
	return &LocationCollection{
		Collection: make(map[uint]*Location),
	}
}

func (lc *LocationCollection) Add(location Location) {
	lc.Collection[location.GetId()] = &location
}

func (lc *LocationCollection) ParseDirections() error {
	for _, location := range lc.Collection {
		log.Println(location)
	}
	return nil
}
