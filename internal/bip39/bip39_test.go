package bip39

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesToBinStr(t *testing.T) {
	expected := "00000100"
	actual := bytesToBinStr([]byte{0x04})
	assert.Equal(t, expected, actual)
}

func TestBinToInt(t *testing.T) {
	expected := 4
	actual, err := binToInt("00000100")
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestChecksum128(t *testing.T) {
	entropy := []byte{0x04}
	expected := "1110"
	actual := checksum(entropy, 128)
	assert.Equal(t, expected, actual)
}

func TestChecksum256(t *testing.T) {
	entropy := []byte{0x04}
	expected := "11100101"
	actual := checksum(entropy, 256)
	assert.Equal(t, expected, actual)
}

func TestGenerateMnemonic128Words(t *testing.T) {
	wordList, err := LoadWordList("../../data/words.txt")
	assert.NoError(t, err)

	mnemonic, err := GenerateMnemonic(128, wordList)
	assert.NoError(t, err)
	words := strings.Fields(mnemonic)
	assert.Equal(t, 12, len(words))

	mnemonic, err = GenerateMnemonic(192, wordList)
	assert.NoError(t, err)
	words = strings.Fields(mnemonic)
	assert.Equal(t, 18, len(words))

	mnemonic, err = GenerateMnemonic(256, wordList)
	assert.NoError(t, err)
	words = strings.Fields(mnemonic)
	assert.Equal(t, 24, len(words))
}
