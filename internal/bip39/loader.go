package bip39

import (
	"bufio"
	"fmt"
	"os"
)

// Returns a list of English words loaded from file.
func LoadWordList(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open word list: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read word list: %w", err)
	}

	return words, nil
}
