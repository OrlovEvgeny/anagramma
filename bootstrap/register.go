package bootstrap

import "context"

type KeyType int

const (
	HashMapAlgorithm KeyType = iota
)

//Register
func Register(ctx context.Context, key KeyType, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}
