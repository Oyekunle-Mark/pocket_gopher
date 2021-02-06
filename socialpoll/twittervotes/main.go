package main

import (
	"gopkg.in/mgo.v2"
	"log"
)

var db *mgo.Session

func main() {}

func dialDb() error {
	var err error
	log.Println("dialing mongodb: localhost")

	db, err = mgo.Dial("localhost")

	return err
}

func closeDb() {
	db.Close()
	log.Println("closed database connection")
}
