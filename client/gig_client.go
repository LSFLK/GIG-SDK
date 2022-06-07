package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"github.com/lsflk/gig-sdk/constants/routes"
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const requestHeaderKey = "User-Agent"
const requestHeaderValue = "Mozilla/5.0 (compatible; Googlebot/2.1; +https://www.google.com/bot.html)"

type Response struct {
	Status  int    `json:"status"`
	Payload string `json:"payload"`
}

type GigClient struct {
	ApiUrl                 string `json:"api_url" bson:"api_url"`
	ApiKey                 string `json:"api_key" bson:"api_key"`
	NerServerUrl           string `json:"ner_server_url" bson:"ner_server_url"`
	NormalizationServerUrl string `json:"normalization_url" bson:"normalization_url"`
	OcrServerUrl           string `json:"ocr_server_url" bson:"ocr_server_url"`
}

/*
GetRequest - get the response string for a given url
*/
func (c GigClient) GetRequest(uri string) (string, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := http.Client{Transport: transport, Timeout: 30 * time.Second}
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set(requestHeaderKey, requestHeaderValue)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, bodyError := ioutil.ReadAll(resp.Body)
	if bodyError != nil {
		return "", bodyError
	}
	err = resp.Body.Close()
	if err != nil {
		return "", err
	}

	return string(body), nil
}

/*
PostRequest - Post to an url with data
*/
func (c *GigClient) PostRequest(uri string, data interface{}) (string, error) {

	// json encode interface
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ApiKey", "ApiKey "+c.ApiKey)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, bodyError := ioutil.ReadAll(resp.Body)
	if bodyError != nil {
		return "", bodyError
	}
	err = resp.Body.Close()
	if err != nil {
		return "", err
	}
	return string(body), err
}

/*
GetEntity - get entity by title
*/
func (c *GigClient) GetEntity(title string) (models.Entity, error) {
	var entity models.Entity
	resp, err := c.GetRequest(c.ApiUrl + routes.Entity + title)
	if err != nil {
		return entity, err
	}
	err = json.Unmarshal([]byte(resp), &entity)
	if err != nil {
		return entity, err
	}
	return entity, err
}

/*
CreateEntityFromText - create entity from text content
*/
func (c *GigClient) CreateEntityFromText(textContent string, title string, categories []string, entityTitles []models.NERResult) error {
	//decode to entity
	var entities []models.Entity
	entity := *new(models.Entity).SetTitle(*new(models.Value).SetType(ValueType.String).SetValueString(title)).
		SetAttribute("content", *new(models.Value).
			SetType(ValueType.String).
			SetValueString(textContent)).
		AddCategories(categories)

	for _, entityObject := range entityTitles {
		//normalizedName, err := utils.NormalizeName(entityObject.EntityName)
		//if err == nil {
		childEntity := *new(models.Entity).
			SetTitle(*new(models.Value).SetType(ValueType.String).
				SetValueString(entityObject.EntityName)).
			AddCategory(entityObject.Category)
		entities = append(entities, childEntity)
		//}
	}

	err := c.AddEntitiesAsLinks(&entity, entities)
	if err != nil {
		panic(err)
	}

	//save to db
	entity, saveErr := c.CreateEntity(entity)

	return saveErr
}

/*
AddEntityAsAttribute - Add entity as an attribute to a given entity
*/
func (c *GigClient) AddEntityAsAttribute(entity models.Entity, attributeName string, attributeEntity models.Entity) (models.Entity, models.Entity, error) {
	entity, linkEntity, err := c.AddEntityAsLink(entity, attributeEntity)
	if err != nil {
		return entity, attributeEntity, err
	}
	entity.SetAttribute(attributeName, models.Value{
		ValueType:   ValueType.String,
		ValueString: linkEntity.Title,
	})

	return entity, linkEntity, nil
}

/*
AddEntityAsLink - Add entity as a link to a given entity
*/
func (c *GigClient) AddEntityAsLink(entity models.Entity, linkEntity models.Entity) (models.Entity, models.Entity, error) {
	createdLinkEntity, linkEntityCreateError := c.CreateEntity(linkEntity)
	if linkEntityCreateError != nil {
		return entity, createdLinkEntity, linkEntityCreateError
	}
	entity.AddLink(*new(models.Link).SetTitle(linkEntity.GetTitle()).AddDate(entity.GetSourceDate()))
	return entity, createdLinkEntity, nil
}

/*
AddEntitiesAsLinks - Add list of related entities to a given entity
*/
func (c *GigClient) AddEntitiesAsLinks(entity *models.Entity, linkEntities []models.Entity) error {
	createdLinkEntities, linkEntityCreateError := c.CreateEntities(linkEntities)
	if linkEntityCreateError != nil {
		return linkEntityCreateError
	}
	for _, linkEntity := range createdLinkEntities {
		newLink := models.Link{}
		newLink.SetTitle(linkEntity.GetTitle()).AddDate(entity.GetSourceDate())
		entity.AddLink(newLink)
	}
	return nil
}

/*
CreateEntity - Create a new entity and save to GIG
*/
func (c *GigClient) CreateEntity(entity models.Entity) (models.Entity, error) {

	resp, err := c.PostRequest(c.ApiUrl+routes.Add, entity)
	if err != nil {
		return entity, err
	}
	err = json.Unmarshal([]byte(resp), &entity)
	if err != nil {
		return entity, err
	}
	return entity, err
}

/*
CreateEntities - Create a list of new entities and save to GIG
*/
func (c *GigClient) CreateEntities(entities []models.Entity) ([]models.Entity, error) {

	_, err := c.PostRequest(c.ApiUrl+routes.AddBatch, entities)
	if err != nil {
		return entities, err
	}
	return entities, err
}

/*
ExtractEntityNames - NER extraction
*/
func (c *GigClient) ExtractEntityNames(textContent string) ([]models.NERResult, error) {

	apiResp, err := c.PostRequest(c.NerServerUrl, textContent)
	if err != nil {
		return nil, err
	}
	var (
		entityTitles [][]string
		results      []models.NERResult
	)
	if err = json.Unmarshal([]byte(apiResp), &entityTitles); err != nil {
		return nil, err
	}

	for _, entity := range entityTitles {
		newNERResult := *new(models.NERResult).SetCategory(entity[1]).SetEntityName(entity[0])
		results = append(results, newNERResult)
	}
	return results, nil
}

/*
NormalizeName - normalize entity title before appending
*/
func (c *GigClient) NormalizeName(title string) (string, error) {

	response, err := c.GetRequest(c.NormalizationServerUrl + routes.Normalize + "?searchText=" + url.QueryEscape(title))
	if err != nil {
		return "", err
	}
	var result Response
	if json.Unmarshal([]byte(response), &result); err != nil {
		return "", err
	}
	log.Println("normalized", title, "to", result.Payload)
	return result.Payload, nil
}

/*
UploadFile - Upload an image through API
*/
func (c *GigClient) UploadFile(payload models.Upload) error {

	if _, err := c.PostRequest(c.ApiUrl+routes.Upload, payload); err != nil {
		return err
	}
	return nil
}
