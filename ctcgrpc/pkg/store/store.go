package store

import (
	"context"
	"fmt"
	"github.com/pradlerfilip/CTC_2022/ctcgrpc/pkg"
	v3 "go.etcd.io/etcd/client/v3"
)

type Store interface {
	Get(ctx context.Context, key string) (string, error)
	Put(ctx context.Context, key, value string) error
	Delete(ctx context.Context, key string) error
}

type etcd struct {
	cl *v3.Client
}

var _ = Store(&etcd{})

func NewEtcd(cl *v3.Client) *etcd {
	return &etcd{cl: cl}
}

func (e *etcd) Get(ctx context.Context, key string) (string, error) {
	r, err := e.cl.Get(ctx, key)

	if err != nil {
		return "", fmt.Errorf("key %v: %w", key, err)
	}

	if len(r.Kvs) != 1 {
		return "", pkg.ErrNotFound
	}

	return string(r.Kvs[0].Value), nil
}

func (e *etcd) Put(ctx context.Context, key, value string) error {
	_, err := e.cl.Put(ctx, key, value)
	return err
}

func (e *etcd) Delete(ctx context.Context, key string) error {
	_, err := e.cl.Delete(ctx, key)
	return err
}
