package storage

/*
Test functions for database abstractions
Test are performed on a seperate database with same traits.
*/

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetConstraints(t *testing.T) {

	a0 := Arm{0, 0, 0}
	a1 := Arm{0, 1, 1}
	a2 := Arm{0, 2, 2}
	a3 := Arm{0, 3, 3}
	tcp := Arm{0, 4, 4}

	testCases := []struct {
		description     string
		elements        []Arm
		expected_result bool
	}{
		{"Correct", []Arm{a0, a1, a2, a3, tcp}, true},
		{"Zero Elements,", []Arm{}, false},
		{"Duplicated Elements", []Arm{a0, a1, a1, a2, a3, tcp}, false},
		{"Missing Elements", []Arm{a0, a2, a3, tcp}, false},
	}

	db, err := gorm.Open(sqlite.Open("TestGetConstraints.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schemas
	db.AutoMigrate(&Constraints{})

	// Insert into table
	for _, v := range testCases {
		for _, x := range v.elements {
			err := AddConstraint(x)
			if err != nil {
				t.Fatalf("Insert of %v ailed", x)
			}
		}
		// Do test

	}

	// Remove elements from table

	// Are zero elements detected?

	// Are duplicates found?

	//Are missing elements detected?
}
