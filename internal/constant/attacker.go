package constant

type AttackerState uint8

const (
	Dead AttackerState = iota
	Trapped
	Attacking
)
