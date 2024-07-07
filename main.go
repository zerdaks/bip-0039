package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	wordList := LoadWordList()
	mnemonic := GenerateMnemonic(128, wordList)
	words := strings.Fields(mnemonic)
	fmt.Printf("Mnemonic with 128 bits of entropy (%d words): ", len(words))
	fmt.Println(mnemonic)
}

// Returns a list of English words loaded from file.
func LoadWordList() []string {
	file, err := os.Open("english.txt")
	if err != nil {
		panic(fmt.Sprintf("error opening file: %v", err))
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("error scanning file: %v", err))
	}
	return lines
}
