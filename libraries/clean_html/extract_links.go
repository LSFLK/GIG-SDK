package clean_html

import (
	"github.com/lsflk/gig-sdk/libraries"
	"github.com/lsflk/gig-sdk/models"
	"golang.org/x/net/html"
	"time"
)

func (c HtmlCleaner) extractLinks(startTag string, n *html.Node, uri string, linkedEntities []models.Entity) (string, []models.Entity) {
	if n.Data == "a" {
		var (
			href  html.Attribute
			title html.Attribute
		)
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				href = attr
			} else if attr.Key == "title" {
				title = attr
			}
		}
		fixedURL := libraries.FixUrl(href.Val, uri)
		if libraries.ExtractDomain(uri) == "en.wikipedia.org" &&
			len(href.Val) > 0 &&
			string(href.Val[0]) != "#" &&
			title.Val != "" &&
			!libraries.StringContainsAnyInSlice(c.Config.IgnoreTitles, title.Val) {
			linkedEntity := models.Entity{Source: fixedURL}
			linkedEntity.SetTitle(models.Value{
				ValueType:   "string",
				ValueString: title.Val,
				Date:        time.Now(),
			})
			linkedEntities = append(linkedEntities, linkedEntity)

		}
		startTag = n.Data + " href='" + fixedURL + "' title='" + title.Val + "'"
	}

	return startTag, linkedEntities
}
