package main

import (
	"fmt"
	"strings"

	"github.com/zerdaks/bip-0039/internal/bip39"
)

func main() {
	wordList := bip39.LoadWordList("data/words.txt")
	mnemonic := bip39.GenerateMnemonic(128, wordList)
	words := strings.Fields(mnemonic)
	fmt.Printf("Mnemonic with 128 bits of entropy (%d words): ", len(words))
	fmt.Println(mnemonic)
}
