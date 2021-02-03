package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	duplicateVowel bool = true
	removeVowel    bool = false
)

func randBool() bool {
	return rand.Intn(2) == 0
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		word := []byte(s.Text())

		if randBool() {
			var vowelIndex int = -1

			for index, char := range word {
				switch char {
				case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
					if randBool() {
						vowelIndex = index
					}
				}
			}

			if vowelIndex >= 0 {
				switch randBool() {
				case duplicateVowel:
					word = append(word[:vowelIndex+1], word[vowelIndex:]...)
				case removeVowel:
					word = append(word[:vowelIndex], word[vowelIndex+1:]...)
				}
			}
		}

		fmt.Println(string(word))
	}
}
