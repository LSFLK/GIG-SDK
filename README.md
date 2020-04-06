# Documentation of GIG SDK
Models, Enums, Tools, Libraries, and API call examples for GIG Development

See [Libraries](libraries/README.md) for generic functions.

See [Crawlers](pdf_crawler/README.md) for sample PDF crawler.

See [Request Handlers](request_handlers/README.md) for sample API calls.

## Configuration
Change GIG server urls in config.go

    var ApiUrl = "http://localhost:9000/api/"
    var NERServerUrl = "http://127.0.0.1:8080/classify"
    var NormalizeServer = "http://localhost:9000/api/normalize"