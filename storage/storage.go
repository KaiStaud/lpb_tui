/*
* Storage functions for interfacing with SQLite database.
* Modul handles operations on robot physical configuration,
* Accessing and changing profiles
 */

package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Coordinates struct {
	gorm.Model
	Name string
	X    int
	Y    int
	Z    int
}

type Configuration struct {
	gorm.Model
	size     float64
	position int
}

type Constraints struct {
	gorm.Model
	ROM_min float64
}

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
	var product Coordinates
	db.First(&product, 1) // find product with integer primary key
}
