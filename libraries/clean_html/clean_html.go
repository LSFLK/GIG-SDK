package clean_html

import (
	"github.com/lsflk/gig-sdk/libraries"
	"github.com/lsflk/gig-sdk/models"
	"golang.org/x/net/html"
	"strings"
)

var defaultIgnoreElements = []string{"noscript", "script", "style", "input", "form", "br", "hr"}

type Config struct {
	LineBreakers   []string
	IgnoreElements []string
	IgnoreStrings  []string
	IgnoreTitles   []string
	IgnoreClasses  []string
}

type HtmlCleaner struct {
	Config Config
}

func (c HtmlCleaner) CleanHTML(uri string, body *html.Node) (string, []models.Entity, []models.Upload, string) {
	var (
		result             string
		linkedEntities     []models.Entity
		f                  func(*html.Node)
		imageList          []models.Upload
		defaultImageSource string
		defaultImageArea   int
	)

	lineBreakers := c.Config.LineBreakers
	ignoreElements := append(c.Config.IgnoreElements, defaultIgnoreElements...)
	ignoreStrings := c.Config.IgnoreStrings
	ignoreClasses := c.Config.IgnoreClasses

	f = func(n *html.Node) {
		if !libraries.StringInSlice(ignoreElements, n.Data) && // ignore elements
			!libraries.StringContainsAnyInSlice(ignoreClasses, ExtractClass(n)) { // ignore classes

			endTag := ""
			trimmedData := strings.TrimSpace(n.Data)
			if n.Type == html.TextNode && trimmedData != "" {
				result = result + getNodeContent(n, ignoreStrings, trimmedData)
			} else if n.Type == html.ElementNode {
				startTag := ""
				imageSource := ""
				imageArea := 0
				startTag, linkedEntities = c.extractLinks(startTag, n, uri, linkedEntities)
				startTag, imageList, imageSource, imageArea = ExtractImages(startTag, n, uri, imageList)
				startTag = ExtractIFrames(startTag, n, uri)
				startTag, endTag = getStartAndEndTag(startTag, endTag, n)
				result = result + startTag

				defaultImageArea, defaultImageSource = setDefaultImage(imageSource, imageArea, defaultImageArea, defaultImageSource)
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}

			result = closeElementTag(result, endTag, lineBreakers, n)
		}
	}
	f(body)

	return result, linkedEntities, imageList, defaultImageSource
}

func closeElementTag(result string, endTag string, lineBreakers []string, n *html.Node) string {
	if endTag != "" {
		result = result + endTag
	}
	if libraries.StringInSlice(lineBreakers, n.Data) {
		result = result + "\n"
	}
	return result
}

func getNodeContent(n *html.Node, ignoreStrings []string, trimmedData string) string {
	if !libraries.StringInSlice(ignoreStrings, trimmedData) {
		return n.Data
	}
	return ""
}

func getStartAndEndTag(startTag string, endTag string, n *html.Node) (string, string) {
	if startTag == "" {
		startTag = "<" + n.Data + ">"
	} else {
		startTag = "<" + startTag + ">"
	}
	endTag = "</" + n.Data + ">"

	return startTag, endTag
}

func setDefaultImage(imageSource string, imageArea int, defaultImageArea int, defaultImageSource string) (int, string) {
	//set default image
	if imageSource != "" && (imageArea > defaultImageArea) {
		defaultImageArea = imageArea
		defaultImageSource = imageSource
	}

	return defaultImageArea, defaultImageSource
}
