/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"database/sql"
	"fmt"
	"log"
	"lpb/cmd"

	_ "github.com/go-sql-driver/mysql"
)

type coordinates struct {
	X int
	Y int
	Z int
}

func coordinates_in_dome(cords coordinates) bool {
	return true
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	cmd.Execute()
	//fmt.Println(viper.Get("name"))
	config, err := cmd.LoadConfig()

	if err != nil {
		log.Fatal("error while looading config")
	}
	fmt.Println("Struct:", config)
	fmt.Println("Drivers", sql.Drivers())

	db, err := sql.Open("mysql", "root:passwort@tcp(127.0.0.1:3306)/test") // Just for testing
	if err != nil {
		log.Fatal("Couldn't open DB")
	}
	defer db.Close()

	results, err := db.Query("select * from profiles")
	if err != nil {
		log.Fatal("Error by fetching data", err)
	}
	for results.Next() {
		var (
			id   int
			name string
			x    int
			y    int
			z    int
		)

		err = results.Scan(&id, &name, &x, &y, &z)
		if err != nil {
			log.Fatal("Error when reading data", err)
		}
		fmt.Printf("ID: %d, Name: %s, X: %d, Y: %d, Z: %d", id, name, x, y, z)
	}

}
