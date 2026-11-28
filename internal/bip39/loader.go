package bip39

import (
	"bufio"
	"fmt"
	"io/fs"
	log "log/slog"
	"strings"
)

// LoadWordList returns the words in the named file in fsys, one per line,
// with surrounding whitespace and any byte order mark removed.
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
		words = append(words, strings.TrimSpace(strings.TrimPrefix(scanner.Text(), "\ufeff")))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read word list: %w", err)
	}

	return words, nil
}
