package utils

import (
	"io"
	"net/http"
	"os"
)

func DownloadSystemdUnit(packageName, serviceName string) {

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
