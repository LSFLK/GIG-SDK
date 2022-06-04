# Documentation of GIG SDK
Models, Enums, Tools, Libraries, and API call examples for GIG Development

See [Libraries](libraries/README.md) for generic functions.

## Installation

go get github.com/lsflk/gig-sdk vX.X.X

## Client Configuration
Configure the GigClient as below

    gigClient := GigClient{
    		ApiUrl:                 "http://localhost:9000/api/",
    		ApiKey:                 "[ApiKey]",
    		NerServerUrl:           "http://localhost:8080/classify",
    		NormalizationServerUrl: "http://localhost:9000/api/",
    		OcrServerUrl:           "http://localhost:8081/extract?url=",
    	}

## Build, Test, Publish

* To run test and publish commands you need to set up a local GIG server.
Refer to [https://github.com/LSFLK/GIG] for more details.
* To test NameExtraction and NER related function you need to set up NER server from [https://github.com/lsflk/Standford-NER-python-wrapper]
* To test OCR related functionality setup OCR server from [https://github.com/umayangag/python-tesseract]
* Modify tests/client/config.go file to change the test server url.

### `go mod tidy`
### `./go_test.sh`
### `git tag vX.X.X`
### `git push origin vX.X.X`
### `GOPROXY=proxy.golang.org go list -m github.com/lsflk/gig-sdk@vX.X.X`