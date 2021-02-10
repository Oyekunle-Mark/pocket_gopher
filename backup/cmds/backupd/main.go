package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/matryer/filedb"
	"log"
	"os"
	"os/signal"
	"pocket_gopher/backup"
	"syscall"
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

	var path path

	col.ForEach(func(_ int, data []byte) bool {
		if err := json.Unmarshal(data, &path); err != nil {
			fatalErr = err
			return true
		}

		m.Paths[path.Path] = path.Hash
		return false // carry on
	})

	if fatalErr != nil {
		return
	}

	if len(m.Paths) < 1 {
		fatalErr = errors.New("no paths - use backup tool to add at least one")
		return
	}

	for {
		check(m, col)

		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

		for {
			select {
			case <-time.After(*interval):
				check(m, col)
			case <-signalChan: // stop
				fmt.Println()
				log.Printf("Stopping...")
				return
			}
		}
	}
}

