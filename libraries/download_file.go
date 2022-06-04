package libraries

import (
	"io"
	"net/http"
	"os"
)

/*
DownloadFile - download a file given the source and destination
*/
func DownloadFile(filePath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
