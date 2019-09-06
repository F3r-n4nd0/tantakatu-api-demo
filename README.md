# Tantakatu API Demo


This is a JSON over HTTP service for a flea-market-like called Tantakatu.


## Technology Stack

* The service is written in Go using [Go-Swagger](https://github.com/go-swagger/go-swagger) to generate the server skeleton from a [Swagger Spec](spec).
* multi-file-swagger (a Nodejs package) is used to aggregate the spec to a single JSON file from the many YML files that define the models, routes etc
* Code is hosted on GitHub
* Go dependencies are managed with Gp Modules [Vgo](https://blog.golang.org/using-go-modules)

## Security

* An API Key set as header X-API-KEY 

## Build Instructions

* Create a new go workspace and clone repo
* Install recent version of NodeJs if not already present
* Install multi-file-swagger [Multi-file Swagger](https://github.com/mohsen1/multi-file-swagger-example)
* Install go-swagger [Go Swagger](https://github.com/go-swagger/go-swagger)
* Run [scripts/swagger_code_gen.sh](scripts/swagger_code_gen.sh) to generate the server skeleton and [client](client) package.
* Set the following environment variables
    * LOG_OUTPUT_FORMAT - JSON_ONLINE | JSON_PRETTY | DEBUG (optional - default DEBUG)
    * LOG_OUTPUT_DEBUG - true/false (optional - default false)
    * MOCK_SERVICE - true/false (optional) there are not production code.
    * API_KEY - the API key used by clients to send email
* Run `go run ./cmd/tantakatu-server/main.go`
