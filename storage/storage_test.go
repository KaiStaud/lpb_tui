package storage

/*
Test functions for database abstractions
Test are performed on a seperate database with same traits.
*/

import (
	"errors"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestAddArms(t *testing.T) {

	a0 := Arm{1, 1, 1, 0}
	a1 := Arm{2, 2, 2, 0}
	a2 := Arm{3, 3, 3, 0}
	a3 := Arm{4, 4, 4, 0}
	tcp := Arm{5, 5, 5, 0}

	testCases := []struct {
		description     string
		elements        []Arm
		new_element     Arm
		expected_result error
	}{
		{"Correct", []Arm{a0, a1, a2, a3}, tcp, nil},
		{"Duplicated Elements", []Arm{a0, a1, a2, a3, tcp}, a0, InsertError},
	}

	db, err := gorm.Open(sqlite.Open("testing.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schemas
	db.AutoMigrate(&Arm{})

	// Insert into table
	for _, subtest := range testCases {
		for _, x := range subtest.elements {
			AddArm(x, db)
		}
		// Do test
		t.Run(subtest.description, func(t *testing.T) {
			err := AddArm(subtest.new_element, db)
			if errors.Is(err, subtest.expected_result) {
				t.Log("Constraint checked correct")
			} else {
				t.Errorf("Checking incorrect, expected %v, got %v", subtest.expected_result, err)
			}

		})
		// Remove elements from table
		for _, x := range subtest.elements {
			db.Delete(&x, x.ID)
		}
		db.Delete(&subtest.new_element, subtest.new_element.ID)
	}
}
