package main

import (
	"context"
	"errors"
	openapi "github.com/soichisumi-sandbox/gcp-sandbox/gen/go"
	"net/http"
	"sync"
)

func NewKVSAPI() *kvsApi {
	return &kvsApi{
		m: sync.Map{},
	}
}

type kvsApi struct {
	m sync.Map
}

func (k *kvsApi) KvsKeyGet(ctx context.Context, s string) (openapi.ImplResponse, error) {
	v, ok := k.m.Load(s)
	if !ok {
		return openapi.Response(http.StatusNotFound, nil), errors.New("not found")
	}
	return openapi.Response(http.StatusOK, v), nil
}

func (k *kvsApi) KvsPost(ctx context.Context, pair openapi.KvPair) (openapi.ImplResponse, error) {
	k.m.Store(pair.Key, pair.Value)
	return openapi.Response(http.StatusOK, nil), nil
}
