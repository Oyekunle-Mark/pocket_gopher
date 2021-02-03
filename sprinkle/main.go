package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const swapWordIdentifier = "*"

var transforms = []string{
	swapWordIdentifier,
	swapWordIdentifier + "app",
	swapWordIdentifier + "site",
	swapWordIdentifier + "time",
	"get" + swapWordIdentifier,
	"go" + swapWordIdentifier,
	"lets " + swapWordIdentifier,
	swapWordIdentifier + "hq",
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		pattern := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(pattern, swapWordIdentifier, s.Text(), -1))
	}
}
