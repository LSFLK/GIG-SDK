package main

import (
	"GIG-SDK/libraries"
	"GIG-SDK/request_handlers"
	"flag"
	"log"
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
		log.Println("starting url not specified")
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
			log.Println(link, uri)
			absoluteUrl := libraries.FixUrl(link, uri)
			log.Println(absoluteUrl)

			// make directory if not exist
			if err = libraries.EnsureDirectory(baseDir); err != nil {
				log.Println(err)
			}

			// download file
			encodedFileName := libraries.ExtractFileName(absoluteUrl)
			filePath := baseDir + encodedFileName
			err := libraries.DownloadFile(filePath, absoluteUrl)
			if err != nil {
				log.Println(err)
			}

			fileName, _ := url.QueryUnescape(encodedFileName)
			//parse pdf
			textContent := libraries.ParsePdf(filePath)
			//NER extraction
			entityTitles, err := request_handlers.ExtractEntityNames(textContent)
			if err != nil {
				log.Println(err)
			}
			if err = request_handlers.CreateEntityFromText(textContent, libraries.ExtractDomain(uri)+" - "+fileName, categories, entityTitles); err != nil {
				log.Println(err.Error(), absoluteUrl)
			}

		}
	}

	log.Println("pdf crawling completed")

}
