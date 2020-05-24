package namegenerator

import (
	"github.com/ypeckstadt/name-generator/names"
	"math/rand"
	"time"
)

// Country represents country the the name needs to be generated for
type Country string

const (
	// UnitedStates generates names for the US
	UnitedStates Country = "us"

	// France generates names for France
	France Country = "france"
)


// GetRandomFirstName returns a random first name
func GetRandomFirstName(country Country) string {
	switch country {
	case UnitedStates:
		return getRandomElement(names.UnitedStatesFirstNames)
	case France:
		return getRandomElement(names.FranceFirstNames)
	}

	return ""
}

// GetRandomLastName returns a random last name
func GetRandomLastName(country Country) string {
	switch country {
	case UnitedStates:
		return getRandomElement(names.UnitedStatesLastNames)
	case France:
		return getRandomElement(names.FranceLastNames)
	}

	return ""
}

// GetRandomLastName returns a random first and last name
func GetRandomName(country Country) (string, string) {
	switch country {
	case UnitedStates:
		return getRandomElement(names.UnitedStatesFirstNames), getRandomElement(names.UnitedStatesLastNames)
	case France:
		return getRandomElement(names.FranceFirstNames), getRandomElement(names.FranceLastNames)
	}
	return "", ""
}

func getRandomElement(list interface{}) string {
	rand.Seed(time.Now().UTC().UnixNano())
	randomIndex := rand.Intn(len(list.([]string)))
	return list.([]string)[randomIndex]
}
