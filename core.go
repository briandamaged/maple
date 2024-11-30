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

type InversePairing[K comparable, V comparable] struct {
	Value V
	Keys  []K
}

func InversePairings[K comparable, V comparable](m map[K]V) []InversePairing[K, V] {
	d := EmptyDefaulter(func(v V) *InversePairing[K, V] {
		return &InversePairing[K, V]{
			Value: v,
			Keys:  []K{},
		}
	})

	for k, v := range m {
		ip := d.Get(v)
		ip.Keys = append(ip.Keys, k)
	}

	ips := make([]InversePairing[K, V], 0, len(*d.MapPtr))
	for _, v := range *d.MapPtr {
		ips = append(ips, *v)
	}

	return ips
}
