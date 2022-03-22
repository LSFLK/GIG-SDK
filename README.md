# Documentation of GIG SDK
Models, Enums, Tools, Libraries, and API call examples for GIG Development

See [Libraries](libraries/README.md) for generic functions.

## Client Configuration
Configure the GigClient as below

    gigClient := GigClient{
    		ApiUrl:                 "http://localhost:9000/api/",
    		ApiKey:                 "[ApiKey]",
    		NerServerUrl:           "http://localhost:8080/classify",
    		NormalizationServerUrl: "http://localhost:9000/api/",
    		OcrServerUrl:           "http://localhost:8081/extract?url=",
    	}
