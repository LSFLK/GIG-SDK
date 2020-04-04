package main

import (
	"GIG-SDK/libraries"
	"GIG-SDK/request_handlers"
	"flag"
	"fmt"
	"github.com/JackDanger/collectlinks"
	"net/url"
	"os"
	"strings"
)

/**
config before running
 */
var downloadDir = "scripts/crawlers/cache/"
var categories = []string{""}

/**
input a web url containing pdf files
automatically download and create entities for all of pdf files
 */
func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("starting url not specified")
		os.Exit(1)
	}
	uri := args[0]

	resp, err := request_handlers.GetRequest(uri)

	if err != nil {
		panic(err)
	}

	links := collectlinks.All(strings.NewReader(resp))

	if err = libraries.EnsureDirectory(downloadDir); err != nil {
		panic(err)
	}

	baseDir := downloadDir + libraries.ExtractDomain(uri) + "/"
	for _, link := range links {
		if libraries.FileTypeCheck(link, "pdf") {
			fmt.Println(link, uri)
			absoluteUrl := libraries.FixUrl(link, uri)
			fmt.Println(absoluteUrl)

			// make directory if not exist
			if err = libraries.EnsureDirectory(baseDir); err != nil {
				fmt.Println(err)
			}

			// download file
			encodedFileName := libraries.ExtractFileName(absoluteUrl)
			filePath := baseDir + encodedFileName
			err := libraries.DownloadFile(filePath, absoluteUrl)
			if err != nil {
				fmt.Println(err)
			}

			fileName, _ := url.QueryUnescape(encodedFileName)
			//parse pdf
			textContent := libraries.ParsePdf(filePath)
			//NER extraction
			entityTitles, err := request_handlers.ExtractEntityNames(textContent)
			if err != nil {
				fmt.Println(err)
			}
			if err = request_handlers.CreateEntityFromText(textContent, libraries.ExtractDomain(uri)+" - "+fileName, categories, entityTitles); err != nil {
				fmt.Println(err.Error(), absoluteUrl)
			}

		}
	}

	fmt.Println("pdf crawling completed")

}
