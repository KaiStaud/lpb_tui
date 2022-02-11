/*
* Storage functions for interfacing with SQLite database.
* Modul handles operations on robot physical configuration,
* Accessing and changing profiles
 */

package storage

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var InsertError error = errors.New("Constraint already exists!")

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
	result := db.Find(&arms)
	return arms, result.Error
}

func AddArm(a Arm, db *gorm.DB) error {
	// Check if constraint exists
	constr, _ := GetArms(db)
	constraint_exist := false

	for _, v := range constr {
		if v.SiliconID == a.SiliconID {
			constraint_exist = true
		}
	}

	if constraint_exist == false {
		a.ID++
		db.Create(&a)
		return nil
	} else {
		return InsertError
	}
}
