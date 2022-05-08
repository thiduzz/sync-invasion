package nodes

import "errors"

type AttackerFactory struct {
	nameGeneratorFunc func() string
}

type AttackerFactoryInterface interface {
	Generate(attackerType interface{}, attackerId uint) (*Attacker, error)
}

//NewAttackerFactory Generates a factory, accepts a function to allow
// mocking an undeterministic function that generates attacker names
func NewAttackerFactory(nameGeneratorFunc func() string) *AttackerFactory {
	return &AttackerFactory{nameGeneratorFunc: nameGeneratorFunc}
}

//Generate Generates attackers based on the type provided
//Assumption: It might be interesting to add a different type of attacker that behaves differently
func (ft AttackerFactory) Generate(attackerType interface{}, attackerId uint) (*Attacker, error) {
	switch attackerType.(type) {
	case Attacker:
		// utilize the factory provided function and id
		return NewAttacker(attackerId, ft.nameGeneratorFunc), nil
	default:
		return nil, errors.New("invalid attacker factory type")
	}
}
