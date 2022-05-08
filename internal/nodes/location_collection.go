package nodes

type LocationCollection struct {
	Collection   map[uint]*Location
	referenceMap map[string]uint
}

func NewLocationCollection() *LocationCollection {
	return &LocationCollection{
		Collection: make(map[uint]*Location),
	}
}

func (lc *LocationCollection) Add(location *Location) {
	lc.Collection[location.GetId()] = location
	lc.referenceMap[location.GetName()] = location.GetId()
}

func (lc *LocationCollection) GetById(id uint) *Location {
	return lc.Collection[id]
}

func (lc *LocationCollection) GetByName(name string) *Location {
	if referenceId, exists := lc.referenceMap[name]; exists {
		return lc.Collection[referenceId]
	}
	return nil
}
