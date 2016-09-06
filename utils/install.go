package utils

import (
	"io"
	"net/http"
	"net/url"
	"os"
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

	resp, err := http.Get(rawurl)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	f, err := os.OpenFile(destBin, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, resp.Body)

	return nil
}

func main() {
}
