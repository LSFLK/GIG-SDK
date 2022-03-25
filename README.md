# Documentation of GIG SDK
Models, Enums, Tools, Libraries, and API call examples for GIG Development

See [Libraries](libraries/README.md) for generic functions.

## Installation

go get github.com/lsflk/gig-sdk v0.1.1

## Client Configuration
Configure the GigClient as below

    gigClient := GigClient{
    		ApiUrl:                 "http://localhost:9000/api/",
    		ApiKey:                 "[ApiKey]",
    		NerServerUrl:           "http://localhost:8080/classify",
    		NormalizationServerUrl: "http://localhost:9000/api/",
    		OcrServerUrl:           "http://localhost:8081/extract?url=",
    	}

## Buid, Test, Publish

* To run test and publish commands you need to setup a local GIG server.
Refer to [https://github.com/LSFLK/GIG] for more details.
* To test NameExtraction and NER related function you need to setup NER server from [https://github.com/lsflk/Standford-NER-python-wrapper]
* To test OCR related functionality setup OCR server from [https://github.com/umayangag/python-tesseract]
* Modify tests/client/config.go file to change the test server url.

### `go mod tidy`
### `./go_test.sh`
### `git tag v0.1.2`
### `git push origin v0.1.2`
### `GOPROXY=proxy.golang.org go list -m github.com/lsflk/gig-sdk@v0.1.2`