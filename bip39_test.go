package main

import (
	"strings"
	"testing"
)

func TestBytesToBinStr(t *testing.T) {
	expected := "00000100"
	actual := bytesToBinStr([]byte{0x04})
	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}

func TestBinToInt(t *testing.T) {
	expected := 4
	actual := binToInt("00000100")
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestChecksum128(t *testing.T) {
	entropy := []byte{0x04}
	expected := "1110"
	actual := checksum(entropy, 128)
	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}

func TestChecksum256(t *testing.T) {
	entropy := []byte{0x04}
	expected := "11100101"
	actual := checksum(entropy, 256)
	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}

func TestGenerateMnemonic128Words(t *testing.T) {
	wordList := LoadWordList()

	mnemonic := GenerateMnemonic(128, wordList)
	words := len(strings.Fields(mnemonic))
	if words != 12 {
		t.Errorf("expected 12 words, got %d", words)
	}

	mnemonic = GenerateMnemonic(192, wordList)
	words = len(strings.Fields(mnemonic))
	if words != 18 {
		t.Errorf("expected 18 words, got %d", words)
	}

	mnemonic = GenerateMnemonic(256, wordList)
	words = len(strings.Fields(mnemonic))
	if words != 24 {
		t.Errorf("expected 24 words, got %d", words)
	}
}
