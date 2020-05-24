package namegenerator

import (
	"testing"
)

// United States
func TestGetRandomFirstNameForUnitedStates(t *testing.T) {
	firstName := GetRandomFirstName(UnitedStates)
	if len(firstName) == 0 {
		t.Fatalf("Generated first name should not be empty")
	}
}

func TestGetRandomLastNameForUnitedStates(t *testing.T) {
	lastName := GetRandomLastName(UnitedStates)
	if len(lastName) == 0 {
		t.Fatalf("Generated last name should not be empty")
	}
}

func TestGetRandomNameForUnitedStates(t *testing.T) {
	firstName, lastName := GetRandomName(UnitedStates)
	if len(firstName) == 0 {
		t.Fatalf("Generated first name should not be empty")
	}
	if len(lastName) == 0 {
		t.Fatalf("Generated last name should not be empty")
	}
}

// France
func TestGetRandomFirstNameForFrance(t *testing.T) {
	firstName := GetRandomFirstName(France)
	if len(firstName) == 0 {
		t.Fatalf("Generated first name should not be empty")
	}
}

func TestGetRandomLastNameForFrance(t *testing.T) {
	lastName := GetRandomLastName(France)
	if len(lastName) == 0 {
		t.Fatalf("Generated last name should not be empty")
	}
}

func TestGetRandomNameForFrance(t *testing.T) {
	firstName, lastName := GetRandomName(France)
	if len(firstName) == 0 {
		t.Fatalf("Generated first name should not be empty")
	}
	if len(lastName) == 0 {
		t.Fatalf("Generated last name should not be empty")
	}
}

