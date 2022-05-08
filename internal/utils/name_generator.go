package utils

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GenerateRandomName() string {
	caseType := cases.Title(language.English)
	return fmt.Sprintf("%s %s", caseType.String(gofakeit.AdjectiveDescriptive()), gofakeit.PetName())
}
