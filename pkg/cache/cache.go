package cache

import (
	"context"
	"errors"
)

var (
	ErrCacheMiss = errors.New("cache miss")
)

type Provider interface {
	Set(ctx context.Context, key string, value []byte) error
	Get(ctx context.Context, key string) ([]byte, error)
	Del(ctx context.Context, key string) error
}

type SerializationType uint8

const (
	SerializationTypeUnknown SerializationType = iota
	SerializationTypeJson
)

type ObjectCacher[T any] struct {
	provider          Provider
	serializationType SerializationType
}

const KeyPrefix = "LINK_SERVICE"

func createKey(k string) string {
	return KeyPrefix + "." + k
}

func NewObjectCacher[T any](p Provider, st SerializationType) *ObjectCacher[T] {
	return &ObjectCacher[T]{
		provider:          p,
		serializationType: st,
	}
}

func NewJsonObjectCacher[T any](p Provider) *ObjectCacher[T] {
	return NewObjectCacher[T](p, SerializationTypeJson)
}

func (c *ObjectCacher[T]) Set(ctx context.Context, key string, value T) error {
	data, err := c.marshal(value)
	if err != nil {
		return err
	}

	return c.provider.Set(ctx, createKey(key), data)
}

func (c *ObjectCacher[T]) Get(ctx context.Context, key string) (T, error) {
	var t T
	data, err := c.provider.Get(ctx, createKey(key))
	if err != nil {
		if errors.Is(err, ErrCacheMiss) {
			return t, nil
		}
		return t, err
	}
	return t, c.unmarshal(data, &t)

}

func (c *ObjectCacher[T]) Del(ctx context.Context, key string) error {
	return c.provider.Del(ctx, createKey(key))
}
