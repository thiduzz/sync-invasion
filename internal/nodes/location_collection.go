package nodes

type LocationCollection struct {
	collection map[uint]*Location
}

func NewLocationCollection() *LocationCollection {
	return &LocationCollection{
		collection: make(map[uint]*Location),
	}
}

func (l *LocationCollection) Add(location Location) {
	l.collection[location.GetId()] = &location
}
