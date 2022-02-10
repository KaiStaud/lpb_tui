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
	gorm.Model
	ROM_min       float64
	ID            int
	NumberInChain int
}

type Arm struct {
	ROM_min       float64
	ID            int
	NumberInChain int
}

var db *gorm.DB

func Init() {
	db, err := gorm.Open(sqlite.Open("profiles.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schemas
	db.AutoMigrate(&Coordinates{}, Configuration{}, &Constraints{})

	// Create defsult configuration and test profile
	db.Create(&Coordinates{Name: "Test", X: 40, Y: 40})
	db.Create(&Configuration{size: 35, position: 1})
	db.Create(&Configuration{size: 35, position: 2})
	db.Create(&Constraints{ROM_min: 2})
	// Read

}

func Add(name string, x float64, y float64, z float64) error {
	db.Create(&Coordinates{Name: name, X: x, Y: y})
	return nil
}

func Get(id int) (Coordinates, error) {
	var product Coordinates
	db.First(&product, 1) // find product with integer primary key
	return product, nil
}

func GetConstraints() ([]Constraints, error) {
	var constraints []Constraints
	matches := 0
	result := db.Find(&constraints)

	// Check if there are missing arms
	for i := 0; i < int(result.RowsAffected); i++ {
		for _, v := range constraints {
			if v.ID == i {
				matches++
			}
		}
	}
	// The robot is set up correctly,when each db object corresponds to one arm
	if result.RowsAffected == int64(matches) {
		return constraints, result.Error
	} else {
		return nil, errors.New("Arms missing/duplicated!")
	}

}

func AddConstraint(a Arm) error {
	// Check if constraint exists
	constr, err := GetConstraints()
	constraint_exist := false
	for _, v := range constr {
		if reflect.DeepEqual(v, a) {
			constraint_exist = true
		}
	}
	if err == nil && constraint_exist == false {
		db.Create(&Constraints{ROM_min: a.ROM_min, ID: a.ID, NumberInChain: a.NumberInChain})
	} else {
		return errors.New("Constraint already exists!")
	}
	return nil
}
