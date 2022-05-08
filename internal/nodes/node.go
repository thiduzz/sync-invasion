package nodes

import "fmt"

type NodeInterface interface {
	GetId() uint
	GetName() string
	fmt.Stringer
}
