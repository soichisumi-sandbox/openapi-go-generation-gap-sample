/*
 * sample API
 *
 * sample description.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
)

// DefaultApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// KvsKeyGet - 
func (s *DefaultApiService) KvsKeyGet(ctx context.Context, key string) (ImplResponse, error) {
	// TODO - update KvsKeyGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, KvPair{}) or use other options such as http.Ok ...
	//return Response(200, KvPair{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("KvsKeyGet method not implemented")
}

// KvsPost - 
func (s *DefaultApiService) KvsPost(ctx context.Context, kvPair KvPair) (ImplResponse, error) {
	// TODO - update KvsPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(200, map[string]interface{}{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("KvsPost method not implemented")
}
