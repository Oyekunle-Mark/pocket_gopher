package main

import (
	"flag"
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

var fatalError error

func fatal(e error) {
	fmt.Println(e)
	flag.PrintDefaults()
	fatalError = e
}

func main() {
	defer func() {
		if fatalError != nil {
			os.Exit(1)
		}
	}()

	log.Println("Connecting to database...")
	db, err := mgo.Dial("localhost")

	if err != nil {
		fatal(err)
		return
	}

	defer func() {
		log.Println("Closing database connection...")
		db.Close()
	}()

	pollData := db.DB("ballots").C("polls")
}
