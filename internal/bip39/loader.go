package bip39

import (
	"bufio"
	"fmt"
	log "log/slog"
	"os"
)

// Returns a list of English words loaded from file.
func LoadWordList(filepath string) ([]string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open word list: %w", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Error("failed to close file", "error", err)
		}
	}()

	scanner := bufio.NewScanner(f)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read word list: %w", err)
	}

	return words, nil
}
