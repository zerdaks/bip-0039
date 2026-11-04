// Package data embeds the BIP-0039 English word list so it is available
// regardless of the working directory.
package data

import "embed"

//go:embed words.txt
var FS embed.FS
