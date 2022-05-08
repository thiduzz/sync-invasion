package nodes

import "errors"

type AttackerFactory struct {
	nameGeneratorFunc func() string
}

type AttackerFactoryInterface interface {
	Generate(attackerType interface{}, attackerId uint) (*Attacker, error)
}

func NewAttackerFactory(nameGeneratorFunc func() string) *AttackerFactory {
	return &AttackerFactory{nameGeneratorFunc: nameGeneratorFunc}
}

func (ft AttackerFactory) Generate(attackerType interface{}, attackerId uint) (*Attacker, error) {
	switch attackerType.(type) {
	case Attacker:
		return NewAttacker(attackerId, ft.nameGeneratorFunc), nil
	default:
		return nil, errors.New("invalid attacker factory type")
	}
}
