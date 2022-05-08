package nodes

import (
	"errors"
	"github.com/thiduzz/code-kata-invasion/internal/constant"
)

type LocationFactory struct {
	nameGeneratorFunc func() string
}

type LocationFactoryOption struct {
	Key   constant.LocationState
	Value interface{}
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

//Seed Provideds a random collection of locations of size greater or equal 0,
// enable the definition of options that apply for the entire set
func (lf *LocationFactory) Seed(locationQty int, options ...*LocationFactoryOption) *LocationCollection {
	if locationQty < 0 {
		return nil
	}
	collection := NewLocationCollection()
	for id := uint(1); id <= uint(locationQty); id++ {
		location, err := lf.Generate(Location{}, id)
		if err != nil {
			return nil
		}

		if len(options) > 0 {
			for _, option := range options {
				if option.Key == constant.Destroyed {
					location.SetDestroyed(option.Value.(bool))
				}
			}
		}

		collection.Add(location)
	}
	return collection
}
