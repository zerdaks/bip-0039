package bip39

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"strconv"
)

// Generates a mnemonic from a given entropy and word list.
func GenerateMnemonic(entLen int, wordList []string) string {
	// validate entropy length
	if (entLen % 32) != 0 {
		panic("entropy should be multiple of 32 bits")
	}
	if entLen < 128 || entLen > 256 {
		panic("entropy should be between 128 and 256 bits")
	}

	// generate random bytes needed for entropy
	byteLen := entLen / 8
	entropy := make([]byte, byteLen)
	if _, err := io.ReadFull(rand.Reader, entropy[:]); err != nil {
		panic(fmt.Sprintf("error reading random bytes: %v", err))
	}

	// convert the initial entropy to a binary string
	entStr := bytesToBinStr(entropy)

	// generate a checksum
	checksum := checksum(entropy, entLen)

	// append the checksum to the binary string
	entStr = entStr + checksum

	// convert the binary string to a mnemonic
	mnemonic := ""
	for i := 0; i < len(entStr); i += 11 {
		// each 11-bit chunk corresponds to a number between 0 and 2047
		j := binToInt(entStr[i : i+11])
		// each number corresponds to a word in a predefined word list
		mnemonic += wordList[j] + " "
	}
	return mnemonic
}

// Converts a byte slice to a binary string.
func bytesToBinStr(bytes []byte) string {
	out := ""
	for _, b := range bytes {
		out += fmt.Sprintf("%08b", b)
	}
	return out
}

// Generates a checksum by taking the first `ENT/32` bits from the hash of the initial entropy.
func checksum(entropy []byte, entLen int) string {
	hash := sha256.Sum256(entropy)
	hashStr := bytesToBinStr(hash[:])
	checksumLen := entLen / 32
	return hashStr[0:checksumLen]
}

// Converts a binary string to an integer.
func binToInt(str string) int {
	val, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		panic(fmt.Sprintf("error converting binary to integer: %v", err))
	}
	return int(val)
}
