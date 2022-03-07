/*
* Storage functions for interfacing with SQLite database.
* Modul handles operations on robot physical configuration,
* Accessing and changing profiles
 */

package storage

import (
	"errors"

	"github.com/go-gl/mathgl/mgl64"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var InsertError error = errors.New("Constraint already exists!")
var db *gorm.DB

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

func Init(path string) error {
	var err error
	db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schemas
	db.AutoMigrate(&Coordinates{}, Configuration{}, &Constraints{}, &Arm{})
	return nil
}

// Return all saved arms from db
func GetArms() ([]Arm, error) {
	var arms []Arm
	result := db.Find(&arms)
	return arms, result.Error
}

// Add an arm to existing configuration
// Checks if arm is known to lpb's configuration via Silicon ID.
// If unknown, the current configuration needs to be reset.
func AddArm(a Arm) error {
	// Check if constraint exists
	constr, _ := GetArms()
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

// Retrieve TCP vector by name.
// Returns nil and not-null-vector if name is matches exactly once
func GetCoordinatesByName(name string) (error, mgl64.Vec3) {
	var matches []Coordinates
	db.Where("name = ?", name).Find(&matches)

	if len(matches) == 0 {
		return errors.New("No match found"), mgl64.Vec3{}
	} else if len(matches) > 1 {
		return errors.New("Name used multiple times"), mgl64.Vec3{}
	} else {
		return nil, mgl64.Vec3{matches[0].X, matches[0].Y, matches[0].Z}
	}
}

// Retrieve TCP vector by name.
// Returns nil and not-null-vector if name is matches exactly once
func GetIDByName(name string) (error, int) {
	var matches []Coordinates
	db.Where("id = ?", name).Find(&matches)

	if len(matches) == 0 {
		return errors.New("No match found"), 0
	} else if len(matches) > 1 {
		return errors.New("Name used multiple times"), 0
	} else {
		return nil, int(matches[0].ID)
	}
}
