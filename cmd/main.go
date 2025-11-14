package main

import (
	log "log/slog"

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

	log.Info("mnemonic with 128 bits of entropy", "count", len(mnemonic), "mnemonic", mnemonic)
}
