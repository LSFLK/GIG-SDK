package clean_html

import (
	"GIG-SDK/libraries"
	"github.com/lsflk/gig-sdk/models"
	"golang.org/x/net/html"
	"strconv"
)

func ExtractImages(startTag string, n *html.Node, uri string, imageList []models.Upload) (string, []models.Upload, string, int) {
	sourceLink := ""
	var uploadImageClass models.Upload
	imageArea := 0

	if n.Data == "img" {
		var src string

		height, width := 0, 0
		for _, attr := range n.Attr {
			switch attr.Key {
			case "src", "data-src":
				src = setSourceValue(src, attr.Val)
			case "width":
				width = convertStringToIntOrDefault(attr.Val)
			case "height":
				height = convertStringToIntOrDefault(attr.Val)
			}
		}

		sourceLink, uploadImageClass = GenerateImagePath(src, uri)
		imageArea = width * height
		startTag = n.Data + " src='" + sourceLink + "'"
		imageList = append(imageList, uploadImageClass)
	}
	return startTag, imageList, sourceLink, imageArea
}

func setSourceValue(src string, value string) string {
	if src == "" {
		return value
	}
	return src
}

func convertStringToIntOrDefault(stringValue string) int {
	value, err := strconv.Atoi(stringValue)
	if err != nil {
		return 1
	}
	return value
}

func GenerateImagePath(href string, uri string) (string, models.Upload) {
	fixedSrc := libraries.FixUrl(href, uri)
	fileName := libraries.ExtractFileName(fixedSrc)
	bucketName := libraries.ExtractDomain(fixedSrc)
	return "images/" + bucketName + "/" + fileName, models.Upload{Title: bucketName, Source: fixedSrc}
}
