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

type poll struct {
	Options []string
}

func loadOption() ([]string, error) {
	var options []string

	iter := db.DB("ballots").C("polls").Find(nil).Iter()

	var p poll

	for iter.Next(&p) {
		options = append(options, p.Options...)
	}

	iter.Close()
	return options, iter.Err()
}
