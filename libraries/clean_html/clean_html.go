package clean_html

import (
	"GIG-SDK/libraries"
	"GIG-SDK/models"
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
		defaultImageArea  int
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
				if !libraries.StringInSlice(ignoreStrings, trimmedData) {
					result = result + n.Data
				}
			} else if n.Type == html.ElementNode {
				startTag := ""
				imageSource := ""
				imageArea := 0
				startTag, linkedEntities = c.extractLinks(startTag, n, uri, linkedEntities)
				startTag, imageList, imageSource, imageArea = ExtractImages(startTag, n, uri, imageList)

				//set default image
				if imageSource != "" && (imageArea > defaultImageArea) {
					defaultImageArea = imageArea
					defaultImageSource = imageSource
				}

				startTag = ExtractIFrames(startTag, n, uri)

				if startTag == "" {
					result = result + "<" + n.Data + ">"
				} else {
					result = result + "<" + startTag + ">"
				}
				endTag = "</" + n.Data + ">"
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}

			if endTag != "" {
				result = result + endTag
			}
			if libraries.StringInSlice(lineBreakers, n.Data) {
				result = result + "\n"
			}
		}
	}
	f(body)

	return result, linkedEntities, imageList, defaultImageSource
}
