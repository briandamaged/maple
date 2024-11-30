package maple

type Pairing[K comparable, V any] struct {
	Key   K
	Value V
}

func Pairings[K comparable, V any](m map[K]V) []Pairing[K, V] {
	var pairings = make([]Pairing[K, V], 0, len(m))
	for k, v := range m {
		pairings = append(pairings, Pairing[K, V]{
			Key:   k,
			Value: v,
		})
	}
	return pairings
}

type Defaulter[K comparable, V any] struct {
	MapPtr      *map[K]V
	DefaultFunc func(K) V
}

func DefaulterFor[K comparable, V any](m *map[K]V, f func(K) V) Defaulter[K, V] {
	return Defaulter[K, V]{
		MapPtr:      m,
		DefaultFunc: f,
	}
}

func EmptyDefaulter[K comparable, V any](f func(K) V) Defaulter[K, V] {
	m := make(map[K]V)
	return DefaulterFor(&m, f)
}

func (d Defaulter[K, V]) Get(key K) V {
	v, exists := (*d.MapPtr)[key]
	if !exists {
		v = d.DefaultFunc(key)
		(*d.MapPtr)[key] = v
	}
	return v
}
