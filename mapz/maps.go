package mapz

// Keys returns all keys in a map in a none deterministic order
func Keys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// Values returns all values in a map in a none deterministic order
func Values[K comparable, V any](m map[K]V) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

// Equal returns true if all key are present in both maps and map to the same value
func Equal[K, V comparable](m1, m2 map[K]V) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

// EqualFunc returns true if all key are present in both maps and map to the same value, determined by the "eq" func
func EqualFunc[K comparable, V1, V2 any](m1 map[K]V1, m2 map[K]V2, eq func(V1, V2) bool) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || !eq(v1, v2) {
			return false
		}
	}
	return true
}

// Clear will delete all elements from a map
func Clear[K comparable, V any](m map[K]V) {
	for k := range m {
		delete(m, k)
	}
}

// Clone will copy all keys and values of a map in to a new one
func Clone[K comparable, V any](m map[K]V) map[K]V {
	r := make(map[K]V, len(m))
	for k, v := range m {
		r[k] = v
	}
	return r
}

// Copy will copy all entries in src into det
func Copy[K comparable, V any](dst, src map[K]V) {
	for k, v := range src {
		dst[k] = v
	}
}

// DeleteValue will remove all instances where the needle matches a value in the map
func DeleteValue[K comparable, V comparable](m map[K]V, needle V) {
	DeleteFunc(m, func(_ K, v V) bool {
		return needle == v
	})
}

// DeleteFunc will remove all entries from a map where the del function returns true
func DeleteFunc[K comparable, V any](m map[K]V, del func(K, V) bool) {
	for k, v := range m {
		if del(k, v) {
			delete(m, k)
		}
	}
}
