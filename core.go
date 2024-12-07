package maple

// Keys returns a slice containing all the keys from the given map.
// The keys are returned in no particular order.
//
// Type Parameters:
//   - K: the type of the keys in the map, which must be comparable.
//   - V: the type of the values in the map, which can be any type.
//
// Parameters:
//   - m: the map from which to extract the keys.
//
// Returns:
//
//	A slice of keys from the map.
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice containing all the values from the given map.
// The values are returned in no particular order.
//
// K is the type of the map's keys, which must be comparable.
// V is the type of the map's values.
//
// Parameters:
//
//	m - The input map from which to extract the values.
//
// Returns:
//
//	A slice containing all the values from the input map.
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Invert swaps the keys and values of a given map. Since several keys could
// potentially map to the same value, this means that the output maps V to
// a slice of K:
//
//	Invert(map[K]V) --> map[V][]K
//
// The function is generic and works with any types K and V that are comparable.
//
// Example:
//
//	input := map[int]string{1: "a", 2: "b", 3: "a"}
//	output := Invert(input)
//	// output is map[string][]int{"a": {1, 3}, "b": {2}}
//
// Parameters:
//
//	m - The input map to be inverted.
//
// Returns:
//
//	A new map where the keys are the values from the input map, and the values are
//	slices of keys from the input map that had the corresponding value.
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
