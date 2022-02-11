/*
* Storage functions for interfacing with SQLite database.
* Modul handles operations on robot physical configuration,
* Accessing and changing profiles
 */

package storage

import (
	"errors"
	"reflect"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Coordinates struct {
	gorm.Model
	Name string
	X    float64
	Y    float64
	Z    float64
}

type Configuration struct {
	gorm.Model
	size     float64
	position int
}

type Constraints struct {
	ROM_min       float64
	ID            int
	NumberInChain int
}

type Arm struct {
	ID            int
	SiliconID     int
	NumberInChain int
	ROM_min       float64
}

func Init() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("profiles.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schemas
	db.AutoMigrate(&Coordinates{}, Configuration{}, &Constraints{}, &Arm{})
	// Add some arms for testing
	//a0 := Arm{0, 0, 0}
	// a1 := Arm{0, 1, 1}
	// a2 := Arm{0, 2, 2}
	// a3 := Arm{0, 3, 3}
	// tcp := Arm{0, 4, 4}
	// db.Create(&a0)
	// db.Create(&a1)
	// db.Create(&a2)
	// db.Create(&a3)
	// db.Create(&tcp)
	return db, nil
}

func GetArms(db *gorm.DB) ([]Arm, error) {
	var arms []Arm
	matches := 0
	result := db.Find(&arms)

	// Check if there are missing arms
	for i := 0; i < int(result.RowsAffected); i++ {
		for _, v := range arms {
			if v.SiliconID == arms[i].SiliconID {
				matches++
			}
		}
	}
	// The robot is set up correctly,when each db object corresponds to one arm
	if result.RowsAffected == 0 {
		return arms, nil
	}
	if matches == 0 {
		return arms, nil
	} else {
		return nil, errors.New("Arms missing/duplicated!")
	}

}

func AddArm(a Arm, db *gorm.DB) error {
	// Check if constraint exists
	constr, err := GetArms(db)
	constraint_exist := false
	for _, v := range constr {
		if reflect.DeepEqual(v, a) {
			constraint_exist = true
		}
	}
	if err == nil && constraint_exist == false {
		db.Create(&a)
		return nil
	} else {
		return errors.New("Constraint already exists!")
	}
}
