package utils

import (
	"net/url"
	"path"
	"path/filepath"
	"runtime"
)

func getUrl(packageName, binName string) string {
	u, _ := url.Parse(deckServerRoot)
	u.Path = path.Join(u.Path, runtime.GOOS, runtime.GOARCH, packageName, binName)
	return u.String()
}

func InstallBin(packageName, binName string) error {
	// FIXME different package has same binName

	rawurl := getUrl(packageName, binName)

	destBin := filepath.Join(deckLocalBinPath, binName)

	return Download(rawurl, destBin)
}

func main() {
}
