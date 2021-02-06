package main

import (
	"flag"
	"fmt"
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
}
