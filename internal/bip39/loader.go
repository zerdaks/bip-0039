package bip39

import (
	"bufio"
	"fmt"
	"io/fs"
	log "log/slog"
)

// Returns a list of English words loaded from the named file in fsys.
func LoadWordList(fsys fs.FS, name string) ([]string, error) {
	f, err := fsys.Open(name)
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
