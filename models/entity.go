package models

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/libraries"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
	"sort"
	"strings"
	"time"
)

// swagger:model
/*
It is recommended to use get,set functions to access values of the entity.
Directly modify attributes only if you know what you are doing.
*/
type Entity struct {
	Id              bson.ObjectId        `json:"-" bson:"_id,omitempty"`
	Title           string               `json:"title" bson:"title"`
	ImageURL        string               `json:"image_url" bson:"image_url"`
	Source          string               `json:"source" bson:"source"`
	SourceSignature string               `json:"source_signature" bson:"source_signature"`
	SourceDate      time.Time            `json:"source_date" bson:"source_date"`
	Attributes      map[string]Attribute `json:"attributes" bson:"attributes"`
	Links           []Link               `json:"links" bson:"links"`
	Categories      []string             `json:"categories" bson:"categories"`
	CreatedAt       time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at" bson:"updated_at"`
	Snippet         string               `json:"snippet" bson:"snippet"`
	SearchText      string               `json:"search_text" bson:"search_text"`
}

func (e Entity) NewEntity() Entity {
	e.Id = bson.NewObjectId()
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return e
}

func (e Entity) GetId() bson.ObjectId {
	return e.Id
}

func (e *Entity) SetTitle(titleValue Value) *Entity {
	// preprocess title
	title := titleValue.GetValueString()
	title = strings.TrimSpace(strings.NewReplacer(
		"%", "",
		"/", "-",
		"~", "2",
		"?", "",
	).Replace(title))
	titleValue.SetValueString(title)
	e.SetAttribute("titles", titleValue)
	if titleAttribute, err := e.GetAttribute("titles"); err == nil {
		e.Title = titleAttribute.GetValue().GetValueString()
		if e.GetSourceDate().IsZero() && !titleAttribute.GetValue().GetDate().IsZero() {
			e.SourceDate = titleAttribute.GetValue().GetDate()
		}
	}

	return e
}

func (e *Entity) SetNewTitle(newTitle string, source string, sourceDate time.Time) *Entity {
	e.SetAttribute("new_title",
		Value{
			ValueType:   ValueType.String,
			ValueString: newTitle,
			Source:      source,
			Date:        sourceDate,
			UpdatedAt:   time.Now(),
		})
	return e
}

func (e Entity) GetTitle() string {
	return e.Title
}

func (e *Entity) SetImageURL(value string) *Entity {
	e.ImageURL = value
	e.UpdatedAt = time.Now()

	return e
}

func (e *Entity) GetImageURL() string {
	return e.ImageURL
}

func (e *Entity) SetSource(value string) *Entity {
	e.Source = value
	e.UpdatedAt = time.Now()

	return e
}

func (e Entity) GetSource() string {
	return e.Source
}

func (e *Entity) SetSourceSignature(value string) *Entity {
	e.SourceSignature = value
	e.UpdatedAt = time.Now()

	return e
}

func (e Entity) GetSourceSignature() string {
	return e.SourceSignature
}

func (e *Entity) SetSourceDate(value time.Time) *Entity {
	e.SourceDate = value
	e.UpdatedAt = time.Now()

	return e
}

func (e Entity) GetSourceDate() time.Time {
	return e.SourceDate
}

/*
SetAttribute - Add or update an existing attribute with a new value
*/
func (e *Entity) SetAttribute(attributeName string, value Value) *Entity {
	//iterate through all attributes
	value.UpdatedAt = time.Now()
	if e.Attributes == nil {
		e.Attributes = make(map[string]Attribute)
	}
	attribute, attributeFound := e.GetAttribute(attributeName)
	valueDate := value.GetDate()
	valueByDate := attribute.GetValueByDate(valueDate)

	if attributeFound == nil {
		valueIndex := -1
		valuesSlice := attribute.GetValues()

		for i, existingValue := range valuesSlice {
			if existingValue.GetValueString() == value.GetValueString() {
				valueIndex = i
				break
			}
		}
		// if the new value doesn't exist already
		if valueIndex == -1 {
			attribute.SetValue(value)
			e.Attributes[attributeName] = attribute // append new value to the attribute

			// if value exist but the value source date is missing
		} else if !(valueIndex == -1 || valueDate.IsZero()) && valuesSlice[valueIndex].GetDate().IsZero() {
			valuesSlice[valueIndex].SetDate(valueDate).SetSource(value.GetSource())
			attribute.Values = valuesSlice
			e.Attributes[attributeName] = attribute
		} else if valueByDate.GetValueString() != value.GetValueString() {
			attribute.SetValue(value)
			e.Attributes[attributeName] = attribute // append new value to the attribute
		}

	} else { //else create new attribute and append value
		attribute := new(Attribute).SetName(attributeName).SetValue(value)
		e.Attributes[attributeName] = *attribute
	}
	e.UpdatedAt = time.Now()
	return e
}

/*
GetAttribute - Get an attribute
*/
func (e Entity) GetAttribute(attributeName string) (Attribute, error) {
	if attribute, attributeFound := e.Attributes[attributeName]; attributeFound {
		return attribute, nil
	}
	return Attribute{}, errors.New("Attribute not found.")
}

func (e Entity) GetAttributes() map[string]Attribute {
	return e.Attributes
}

func (e *Entity) RemoveAttribute(attributeName string) *Entity {
	delete(e.Attributes, attributeName)
	return e
}

/*
AddLink - Add new link to entity
*/
func (e *Entity) AddLink(link Link) *Entity {
	title, dates := link.GetTitle(), link.GetDates()
	if title == "" {
		return e
	}

	var existingLink Link
	for i, linkItem := range e.GetLinks() {
		if linkItem.GetTitle() == title {
			existingLink = linkItem
			for _, date := range dates {
				existingLink.AddDate(date)
			}
			e.Links[i] = existingLink
			return e
		}
	}
	e.Links = append(e.Links, link)
	return e
}

/*
AddLinks - Add new links to entity
*/
func (e *Entity) AddLinks(links []Link) *Entity {
	entity := e
	for _, link := range links {
		entity = e.AddLink(link)
	}
	return entity
}

func (e Entity) GetLinks() []Link {
	return e.Links
}

func (e Entity) GetLinkTitles() []string {
	var titlesSlice []string
	for _, link := range e.GetLinks() {
		titlesSlice = append(titlesSlice, link.GetTitle())
	}
	sort.Slice(titlesSlice, func(i, j int) bool {
		return titlesSlice[i] < titlesSlice[j]
	})
	return titlesSlice
}

/*
SetSnippet - Create snippet for the entity
*/
func (e *Entity) SetSnippet() *Entity {
	contentAttr, err := e.GetAttribute("content")
	snippet := ""
	if err == nil { // if content attribute found
		switch contentAttr.GetValue().GetType() {
		case "html":
			newsDoc, _ := goquery.NewDocumentFromReader(strings.NewReader(contentAttr.GetValue().GetValueString()))
			snippet = strings.Replace(newsDoc.Text(), "  ", "", -1)
		default:
			snippet = contentAttr.GetValue().GetValueString()
		}
	}
	e.SearchText = snippet
	if e.Snippet == "" {
		if len(snippet) > 300 {
			e.Snippet = snippet[0:300] + "..."
		}
		e.Snippet = snippet
	}
	return e
}

func (e Entity) GetSnippet() string {
	return e.Snippet
}

/*
HasContent - Check if the entity has data
*/
func (e Entity) HasContent() bool {
	if len(e.Links) != 0 {
		return true
	}
	if len(e.Categories) != 0 {
		return true
	}
	if len(e.Attributes) != 0 {
		return true
	}
	return false
}

/*
IsNil - Check if the entity has no title, data
*/
func (e Entity) IsNil() bool {
	if e.GetTitle() != "" {
		return false
	}
	return !e.HasContent()
}

/*
AddCategory - Add new category to entity
*/
func (e *Entity) AddCategory(category string) *Entity {
	if libraries.StringInSlice(e.GetCategories(), category) {
		return e
	}
	e.Categories = append(e.GetCategories(), category)
	e.UpdatedAt = time.Now()
	return e
}

/*
AddCategories - Add new categories to entity
*/
func (e *Entity) AddCategories(categories []string) *Entity {
	for _, category := range categories {
		e.AddCategory(category)
	}
	return e
}

/*
RemoveCategories - remove categories from the entity
*/
func (e *Entity) RemoveCategories(categories []string) *Entity {
	var remainingCategories []string
	for _, category := range e.GetCategories() {
		if !libraries.StringInSlice(categories, category) {
			remainingCategories = append(remainingCategories, category)
		}
	}
	e.Categories = remainingCategories
	return e
}

func (e Entity) GetCategories() []string {
	return e.Categories
}

func (e Entity) GetCreatedDate() time.Time {
	return e.CreatedAt
}

func (e Entity) GetUpdatedDate() time.Time {
	return e.UpdatedAt
}

func (e Entity) IsTerminated() bool {
	return strings.Contains(e.GetTitle(), " - Terminated on ")
}
