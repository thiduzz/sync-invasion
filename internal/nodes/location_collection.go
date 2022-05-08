package nodes

type LocationCollection struct {
	Collection map[uint]*Location
}

func NewLocationCollection() *LocationCollection {
	return &LocationCollection{
		Collection: make(map[uint]*Location),
	}
}

func (l *LocationCollection) Add(location Location) {
	l.Collection[location.GetId()] = &location
}
