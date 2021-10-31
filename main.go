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
	"context"
	"fmt"
	"log"
	"lpb/cmd"
	"lpb/handlers"
	"os"
	"os/signal"
	"time"

	"net/http"

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

	l := log.New(os.Stdout, "profile-api", log.LstdFlags)

	// Create a new serve-mux and register the profile handler
	sm := http.NewServeMux()
	ph := handlers.NewProfiles(l)
	sm.Handle("/", ph) // add profile handler to created serve mux

	// Accept all requests @ port 9090
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Handle http requests in go routine
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Catch OS-Signals, gracefully shutdown the services
	// Open a new channel for this:
	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, os.Kill)

	// Catch OS-Signals
	sig := <-sigchan
	l.Println("Received terminate, gracefully shutting down", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
