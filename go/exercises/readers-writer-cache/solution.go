package readerswritercache

// ReadersWriterCache should use sync.RWMutex and run a missing-key loader once.
type ReadersWriterCache[K comparable, V any] struct{}

func NewReadersWriterCache[K comparable, V any]() *ReadersWriterCache[K, V] {
	return &ReadersWriterCache[K, V]{}
}
func (*ReadersWriterCache[K, V]) Get(K) (V, bool)                   { var zero V; return zero, false }
func (*ReadersWriterCache[K, V]) Put(K, V)                          {}
func (*ReadersWriterCache[K, V]) GetOrLoad(key K, load func(K) V) V { var zero V; return zero }
