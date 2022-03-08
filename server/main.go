package main

import (
	openapi "github.com/soichisumi-sandbox/gcp-sandbox/gen/go"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	APIController := openapi.NewDefaultApiController(NewKVSAPI())

	router := openapi.NewRouter(APIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}