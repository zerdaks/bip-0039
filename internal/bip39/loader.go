package bip39

import (
	"bufio"
	"fmt"
	"os"
)

// Returns a list of English words loaded from file.
func LoadWordList(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(fmt.Sprintf("error opening file: %v", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("error scanning file: %v", err))
	}

	return words
}
