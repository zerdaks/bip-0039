package bip39

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"io/fs"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zerdaks/bip-0039/data"
)

func TestLoadWordList(t *testing.T) {
	wordList, err := LoadWordList(data.FS, "words.txt")
	assert.NoError(t, err)
	require.Len(t, wordList, 2048)
	assert.Equal(t, "abandon", wordList[0])
	assert.Equal(t, "zoo", wordList[2047])
}

func TestLoadWordListMissingFile(t *testing.T) {
	wordList, err := LoadWordList(data.FS, "missing.txt")
	assert.Error(t, err)
	assert.Nil(t, wordList)
}

func TestLoadWordListLineTooLong(t *testing.T) {
	fsys := fstest.MapFS{
		"words.txt": {Data: []byte(strings.Repeat("a", bufio.MaxScanTokenSize+1))},
	}

	wordList, err := LoadWordList(fsys, "words.txt")
	assert.ErrorIs(t, err, bufio.ErrTooLong)
	assert.Nil(t, wordList)
}

// SHA-256 of https://github.com/bitcoin/bips/blob/master/bip-0039/english.txt
func TestWordListIsCanonical(t *testing.T) {
	expected := "2f5eed53a4727b4bf8880d8f3f199efc90e58503646d9ff8eff3a2ed3b24dbda"

	words, err := fs.ReadFile(data.FS, "words.txt")
	assert.NoError(t, err)

	sum := sha256.Sum256(words)
	assert.Equal(t, expected, hex.EncodeToString(sum[:]))
}
