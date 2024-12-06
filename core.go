package maple

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

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

func Invert[K comparable, V comparable](m map[K]V) map[V][]K {
	retval := map[V][]K{}
	for k, v := range m {
		keys, exists := retval[v]
		if !exists {
			keys = []K{}
		}
		retval[v] = append(keys, k)
	}

	return retval
}
