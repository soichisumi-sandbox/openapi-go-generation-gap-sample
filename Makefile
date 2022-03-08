openapi-generate:
	openapi-generator generate -i openapi.yaml -g go-server -o ./gen

build-container:
	go build