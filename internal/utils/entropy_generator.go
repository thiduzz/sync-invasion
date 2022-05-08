package utils

import (
	"math/rand"
	"time"
)

//GenerateRandomizer Generate a randomizer that allow control
// of undeterministic behavior in the engine of the application
func GenerateRandomizer(seed int64) *rand.Rand {
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	source := rand.NewSource(seed)
	return rand.New(source)
}
