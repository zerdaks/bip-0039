package bip39

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadWordList(t *testing.T) {
	wordList, err := LoadWordList("../../data/words.txt")
	assert.NoError(t, err)
	assert.Equal(t, 2048, len(wordList))
	assert.Equal(t, "abandon", wordList[0])
	assert.Equal(t, "zoo", wordList[2047])
}

func TestLoadWordListMissingFile(t *testing.T) {
	wordList, err := LoadWordList("../../data/missing.txt")
	assert.Error(t, err)
	assert.Nil(t, wordList)
}

// SHA-256 of https://github.com/bitcoin/bips/blob/master/bip-0039/english.txt
func TestWordListIsCanonical(t *testing.T) {
	expected := "2f5eed53a4727b4bf8880d8f3f199efc90e58503646d9ff8eff3a2ed3b24dbda"

	data, err := os.ReadFile("../../data/words.txt")
	assert.NoError(t, err)

	sum := sha256.Sum256(data)
	assert.Equal(t, expected, hex.EncodeToString(sum[:]))
}
