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
	actual := binToInt("00000100")
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
	wordList := LoadWordList("../../data/words.txt")

	mnemonic := GenerateMnemonic(128, wordList)
	words := strings.Fields(mnemonic)
	assert.Equal(t, 12, len(words))

	mnemonic = GenerateMnemonic(192, wordList)
	words = strings.Fields(mnemonic)
	assert.Equal(t, 18, len(words))

	mnemonic = GenerateMnemonic(256, wordList)
	words = strings.Fields(mnemonic)
	assert.Equal(t, 24, len(words))
}
