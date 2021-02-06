package main

import (
	"flag"
	"fmt"
	"github.com/bitly/go-nsq"
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"sync"
)

var fatalError error
var counts map[string]int
var countsLock sync.Mutex

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

	log.Println("Connecting to nsq...")
	q, err := nsq.NewConsumer("votes", "counter", nsq.NewConfig())

	if err != nil {
		fatal(err)
		return
	}

	q.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
		countsLock.Lock()
		defer countsLock.Unlock()

		if counts == nil {
			counts = make(map[string]int)
		}

		vote := string(m.Body)
		counts[vote]++

		return nil
	}))
}
