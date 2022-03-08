module github.com/soichisumi-sandbox/gcp-sandbox/server

go 1.16

require (
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/soichisumi-sandbox/gcp-sandbox/gen v0.0.0-00010101000000-000000000000
)

replace github.com/soichisumi-sandbox/gcp-sandbox/gen => ../gen