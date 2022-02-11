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

func TestGetArms(t *testing.T) {

	testerror := errors.New("Arms missing/duplicated!")
	a0 := Arm{1, 1, 1, 0}
	a1 := Arm{2, 2, 2, 0}
	a2 := Arm{3, 3, 3, 0}
	a3 := Arm{4, 4, 4, 0}
	tcp := Arm{5, 5, 5, 0}
	duplicate := Arm{6, 2, 2, 0}

	testCases := []struct {
		description     string
		elements        []Arm
		expected_result error
	}{
		{"Correct", []Arm{a0, a1, a2, a3, tcp}, nil},
		{"Zero Elements,", []Arm{}, nil},
		{"Duplicated Elements", []Arm{a0, a1, a2, a3, tcp, duplicate}, testerror},
		{"Missing Elements", []Arm{a0, a2, a3, tcp}, testerror},
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
			db.Create(&x)
		}
		// Do test
		t.Run(subtest.description, func(t *testing.T) {
			_, err := GetArms(db)
			if err == subtest.expected_result {
				t.Log("Constraint checked correct")
			} else {
				t.Errorf("Checking incorrect, expected %v, got %v", subtest.expected_result, err)
			}

		})
		// Remove elements from table
		for _, x := range subtest.elements {
			db.Delete(&x, x.ID)

		}
	}
}
