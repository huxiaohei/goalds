package visitor

type KVisitor[K any] func(key K) bool
type VVisitor[V any] func(value V) bool
type KVVisitor[K any, V any] func(key K, value V) bool
