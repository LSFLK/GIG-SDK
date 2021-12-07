package clean_html

import (
	"GIG-SDK/libraries"
	"GIG-SDK/models"
	"golang.org/x/net/html"
	"strconv"
)

func ExtractImages(startTag string, n *html.Node, uri string, imageList []models.Upload) (string, []models.Upload, string, int) {
	sourceLink := ""
	var uploadImageClass models.Upload
	imageArea := 0

	if n.Data == "img" {
		var (
			src string
			err error
		)
		height, width := 1, 1
		for _, attr := range n.Attr {
			if attr.Key == "src" && src == "" {
				src = attr.Val
			} else if attr.Key == "data-src" {
				src = attr.Val
			} else if attr.Key == "width" {
				width, err = strconv.Atoi(attr.Val)
				if err != nil {
					width = 1
				}
			} else if attr.Key == "height" {
				height, err = strconv.Atoi(attr.Val)
				if err != nil {
					height = 1
				}
			}
		}

		sourceLink, uploadImageClass = GenerateImagePath(src, uri)
		imageArea = width * height
		startTag = n.Data + " src='" + sourceLink + "'"
		imageList = append(imageList, uploadImageClass)
	}
	return startTag, imageList, sourceLink, imageArea
}

func GenerateImagePath(href string, uri string) (string, models.Upload) {
	fixedSrc := libraries.FixUrl(href, uri)
	fileName := libraries.ExtractFileName(fixedSrc)
	bucketName := libraries.ExtractDomain(fixedSrc)
	return "images/" + bucketName + "/" + fileName, models.Upload{Title: bucketName, Source: fixedSrc}
}
