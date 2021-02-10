package main

import (
	"flag"
	"github.com/matryer/filedb"
	"log"
	"pocket_gopher/backup"
	"time"
)

type path struct {
	Path string
	Hash string
}

func main() {
	var fatalErr error

	defer func() {
		if fatalErr != nil {
			log.Fatalln(fatalErr)
		}
	}()

	var (
		interval = flag.Duration("interval", 10*time.Second, "interval between checks")
		archive  = flag.String("archive", "archive", "path to archive location")
		dbpath   = flag.String("db", "./db", "path to filedb database")
	)

	m := &backup.Monitor{
		Destination: *archive,
		Archiver:    backup.ZIP,
		Paths:       make(map[string]string),
	}

	flag.Parse()

	db, err := filedb.Dial(*dbpath)

	if err != nil {
		fatalErr = err
		return
	}

	defer db.Close()
	col, err := db.C("paths")

	if err != nil {
		fatalErr = err
		return
	}
}
