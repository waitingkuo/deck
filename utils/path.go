package utils

import (
	"os"
)

var deckServerRoot = "http://104.199.175.27/deck"
var deckLocalBinPath = "/usr/local/bin"
var deckLocalSystemUnitPath = "/etc/systemd/system"

func MkdirDeckLocalBinPath() {
	if err := os.MkdirAll(deckLocalBinPath, 0755); err != nil {
		panic(err)
	}
}
