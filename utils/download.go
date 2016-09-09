package utils

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

func DownloadSystemdUnit(packageName, serviceName string) error {
	u, err := url.Parse(deckServerRoot)
	if err != nil {
		return err
	}

	u.Path = path.Join(u.Path, "systemd-units", packageName, serviceName+".service")

	rawurl := u.String()
	destPath := filepath.Join(deckLocalSystemUnitPath, serviceName+".service")

	println("rawurl:", rawurl)
	return Download(rawurl, destPath)
}

func Download(rawurl string, destPath string) error {

	resp, err := http.Get(rawurl)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	f, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, resp.Body)

	return nil
}
