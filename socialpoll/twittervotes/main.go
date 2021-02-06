package main

import (
	"github.com/bitly/go-nsq"
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

func publishVotes(votes <- chan string) <- chan struct{} {
	stopchan := make(chan struct{}, 1)

	pub, _ := nsq.NewProducer("localhost:4150", nsq.NewConfig())

	go func() {
		for vote := range votes {
			pub.Publish("votes", []byte(vote)) // publish vote
		}

		log.Println("Publisher: Stopping")
		pub.Stop()
		log.Println("Publisher: Stopped")
		stopchan <- struct{}{}
	}()

	return stopchan
}
