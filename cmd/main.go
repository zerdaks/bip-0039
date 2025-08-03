package main

import (
	"fmt"
	log "log/slog"
	"strings"

	"github.com/zerdaks/bip-0039/internal/bip39"
)

func main() {
	wordList, err := bip39.LoadWordList("data/words.txt")
	if err != nil {
		log.Error(err.Error())
		return
	}

	mnemonic, err := bip39.GenerateMnemonic(128, wordList)
	if err != nil {
		log.Error(err.Error())
		return
	}

	words := strings.Fields(mnemonic)
	log.Info(fmt.Sprintf("mnemonic with 128 bits of entropy (%d words): %s", len(words), mnemonic))
}
