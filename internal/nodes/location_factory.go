package nodes

import "errors"

type LocationFactory struct {
	nameGeneratorFunc func() string
}

type LocationFactoryInterface interface {
	Generate(locationType interface{}, locationId uint) (*Location, error)
	Seed(int) *LocationCollection
}

//NewLocationFactory Generates a factory, accepts a function to allow
func NewLocationFactory(nameGeneratorFunc func() string) *LocationFactory {
	return &LocationFactory{nameGeneratorFunc: nameGeneratorFunc}
}

//Generate Generates attackers based on the type provided
//Assumption: It might be interesting to add a different type of location that behaves differently
func (lf LocationFactory) Generate(locationType interface{}, attackerId uint) (*Location, error) {
	switch locationType.(type) {
	case Location:
		// utilize the factory provided function and id
		return NewLocation(attackerId, lf.nameGeneratorFunc()), nil
	default:
		return nil, errors.New("invalid attacker factory type")
	}
}

//Seed Provideds a random collection of locations of size greater or equal 0
func (lf LocationFactory) Seed(locationQty int) *LocationCollection {
	if locationQty < 0 {
		return nil
	}
	collection := &LocationCollection{}
	for id := uint(1); id <= uint(locationQty); id++ {
		location, err := lf.Generate(Location{}, id)
		if err != nil {
			return nil
		}
		collection.Add(location)
	}
	return collection
}
