package api

import (
	"context"
	"gitlab.com/ondrej.smola/ctcgrpc/pkg/store"
)

type Client interface {
	store.Store
}

type grpcClient struct {
	c ApiClient
}

var _ = Client(&grpcClient{})

func NewGrpcClient(c ApiClient) *grpcClient {
	return &grpcClient{c: c}
}

func (g *grpcClient) Get(ctx context.Context, key string) (string, error) {
	resp, err := g.c.Get(ctx, &GetRequest{Key: key})
	if err != nil {
		return "", err
	}

	return resp.Value, nil
}

func (g *grpcClient) Put(ctx context.Context, key, value string) error {
	_, err := g.c.Put(ctx, &PutRequest{Key: key, Value: value})
	return err
}

func (g *grpcClient) Delete(ctx context.Context, key string) error {
	panic("implement me")
}
