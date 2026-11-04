package main

import (
	log "log/slog"
	"os"

	"github.com/zerdaks/bip-0039/data"
	"github.com/zerdaks/bip-0039/internal/bip39"
)

func main() {
	wordList, err := bip39.LoadWordList(data.FS, "words.txt")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	mnemonic, err := bip39.GenerateMnemonic(128, wordList)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	log.Info("mnemonic with 128 bits of entropy", "count", len(mnemonic), "mnemonic", mnemonic)
}
