package client

import (
	"github.com/lsflk/gig-sdk/client"
)

/**
Set the GIG server API url here for crawlers
 */

var testClient = client.GigClient{
	ApiUrl:                 "http://localhost:9000/api/",
	ApiKey:                 "[ApiKey]",
	NerServerUrl:           "http://localhost:8081/classify",
	NormalizationServerUrl: "http://localhost:9000/api/",
	OcrServerUrl:           "http://localhost:8082/extract?url=",
}
