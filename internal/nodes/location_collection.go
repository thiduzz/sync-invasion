package nodes

type LocationCollection struct {
	collection map[string]*Location
}

func NewLocationCollection() *LocationCollection {
	return &LocationCollection{
		collection: make(map[string]*Location),
	}
}
