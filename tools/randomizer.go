package tools

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math/rand"
	"time"
)

//Randomizer Embeds rand.Rand and enable control over entropy
type Randomizer struct {
	*rand.Rand
	preventReseed bool
}

//NewRandomizer Generate a randomizer that allow control
// of undeterministic behavior in the engine of the application
func NewRandomizer(seed int64) *Randomizer {
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	source := rand.NewSource(seed)
	return &Randomizer{rand.New(source), false}
}

func (rm *Randomizer) PreventReseed(preventReseed bool) {
	rm.preventReseed = preventReseed
}

func (rm *Randomizer) RandomName() string {
	caseType := cases.Title(language.English)
	return fmt.Sprintf("%s %s", caseType.String(gofakeit.AdjectiveDescriptive()), gofakeit.PetName())
}

func (rm *Randomizer) ShuffleUint(slice []uint) {
	if rm == nil || rm.preventReseed {
		return
	}
	rm.Rand.Shuffle(len(slice), func(i, j int) {
		slice[uint(i)], slice[uint(i)] = slice[uint(i)], slice[uint(i)]
	})
}

//Reseed Ensure that randomizer has a different entropy
func (rm *Randomizer) Reseed() {
	if rm == nil || rm.preventReseed {
		return
	}
	rm.Seed(time.Now().UnixNano())
}
