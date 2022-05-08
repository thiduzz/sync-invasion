package constant

type LocationState uint8

const (
	DirectionNorth = "north"
	DirectionSouth = "south"
	DirectionEast  = "east"
	DirectionWest  = "west"
)

const (
	Destroyed LocationState = iota
)
