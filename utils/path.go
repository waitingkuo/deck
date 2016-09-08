package utils

import (
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

var deckServerRoot = "http://104.199.175.27/deck"
var deckLocalBinPath = ""

func init() {
	homeDir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	// consider using /usr/local/bin
	deckLocalRootPath := filepath.Join(homeDir, ".deck")
	deckLocalBinPath = filepath.Join(deckLocalRootPath, "bin")
}

func MkdirDeckLocalBinPath() {
	if err := os.MkdirAll(deckLocalBinPath, 0755); err != nil {
		panic(err)
	}
}
