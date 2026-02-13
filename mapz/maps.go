package mapz

import "github.com/modfin/henry/slicez"

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

// EqualBy returns true if all key are present in both maps and map to the same value, determined by the "eq" func
func EqualBy[K comparable, V1, V2 any](m1 map[K]V1, m2 map[K]V2, eq func(V1, V2) bool) bool {
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
// Warning mutates map
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
// Warning mutates map
func Copy[K comparable, V any](dst, src map[K]V) {
	for k, v := range src {
		dst[k] = v
	}
}

// Merge multiple maps from left to right into a new map.
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	// Pre-calculate capacity to avoid reallocations
	capacity := 0
	for _, m := range maps {
		capacity += len(m)
	}

	out := make(map[K]V, capacity)
	for _, m := range maps {
		for k, v := range m {
			out[k] = v
		}
	}
	return out
}

// DeleteValues will remove all instances where the needle matches a value in the map
// Warning mutates map, use Filter or Reject for immutable version
func DeleteValues[K comparable, V comparable](m map[K]V, needles ...V) {
	set := slicez.Set(needles)
	Delete(m, func(_ K, v V) bool {
		return set[v]
	})
}

// DeleteKeys will remove all instances where the needles matches a key in the map
// Warning mutates map, use Filter or Reject for immutable version
func DeleteKeys[K comparable, V any](m map[K]V, needles ...K) {
	for _, needle := range needles {
		delete(m, needle)
	}
}

// Delete will remove all entries from a map where the del function returns true
// Warning mutates map, , use Filter or Reject for immutable version
func Delete[K comparable, V any](m map[K]V, del func(K, V) bool) {
	for k, v := range m {
		if del(k, v) {
			delete(m, k)
		}
	}
}

func ValueOr[K comparable, V any](m map[K]V, key K, fallback V) V {
	value, exist := m[key]
	if !exist {
		value = fallback
	}
	return value
}

func Filter[K comparable, V any](m map[K]V, pick func(key K, val V) bool) map[K]V {
	// Pre-allocate with capacity of input map
	// In worst case, all elements pass the filter
	res := make(map[K]V, len(m))
	for k, v := range m {
		if pick(k, v) {
			res[k] = v
		}
	}
	return res
}

func FilterByKeys[K comparable, V any](m map[K]V, keys []K) map[K]V {
	set := slicez.Set(keys)
	return Filter(m, func(key K, _ V) bool {
		return set[key]
	})
}
func FilterByValues[K comparable, V comparable](m map[K]V, values []V) map[K]V {
	set := slicez.Set(values)
	return Filter(m, func(_ K, val V) bool {
		return set[val]
	})
}

func Reject[K comparable, V any](m map[K]V, omit func(key K, val V) bool) map[K]V {
	return Filter(m, func(key K, val V) bool {
		return !omit(key, val)
	})
}

func RejectByKeys[K comparable, V any](m map[K]V, keys []K) map[K]V {
	set := slicez.Set(keys)
	return Reject(m, func(key K, _ V) bool {
		return set[key]
	})
}
func RejectByValues[K comparable, V comparable](m map[K]V, values []V) map[K]V {
	set := slicez.Set(values)
	return Reject(m, func(_ K, val V) bool {
		return set[val]
	})
}

func Slice[E any, K comparable, V any](m map[K]V, zip func(K, V) E) []E {
	res := make([]E, 0, len(m))
	for k, v := range m {
		res = append(res, zip(k, v))
	}
	return res
}

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func Entries[K comparable, V any](m map[K]V) []Entry[K, V] {
	return Slice(m, func(k K, v V) Entry[K, V] {
		return Entry[K, V]{k, v}
	})
}
func FromEntries[K comparable, V any](slice []Entry[K, V]) map[K]V {
	return slicez.Associate(slice, func(a Entry[K, V]) (key K, value V) {
		return a.Key, a.Value
	})
}

// Remap manipulates a map keys and values and transforms it to a map of another types.
func Remap[K comparable, V any, K2 comparable, V2 any](in map[K]V, mapper func(K, V) (K2, V2)) map[K2]V2 {
	result := make(map[K2]V2, len(in))

	for k, v := range in {
		k2, v2 := mapper(k, v)
		result[k2] = v2
	}

	return result
}

// RemapKeys manipulates a map keys and transforms it to a map of another types.
func RemapKeys[K comparable, V any, K2 comparable](in map[K]V, mapper func(K, V) K2) map[K2]V {
	return Remap(in, func(k K, v V) (K2, V) {
		return mapper(k, v), v
	})
}

// RemapValues manipulates a map keys and transforms it to a map of another types.
func RemapValues[K comparable, V any, V2 any](in map[K]V, mapper func(K, V) V2) map[K]V2 {
	return Remap(in, func(k K, v V) (K, V2) {
		return k, mapper(k, v)
	})
}

func Invert[K, V comparable](m map[K]V) map[V]K {
	return Remap(m, func(k K, v V) (V, K) {
		return v, k
	})
}

// MapKeys transforms the keys of a map while keeping the same values.
// The mapper function receives only the key (unlike RemapKeys which receives both key and value).
//
// Example:
//
//	MapKeys(map[string]int{"a": 1, "b": 2}, strings.ToUpper)
//	// Returns map[string]int{"A": 1, "B": 2}
func MapKeys[K1 comparable, V any, K2 comparable](m map[K1]V, mapper func(K1) K2) map[K2]V {
	result := make(map[K2]V, len(m))
	for k, v := range m {
		result[mapper(k)] = v
	}
	return result
}

// MapValues transforms the values of a map while keeping the same keys.
// The mapper function receives only the value (unlike RemapValues which receives both key and value).
//
// Example:
//
//	MapValues(map[string]int{"a": 1, "b": 2}, func(v int) int { return v * 2 })
//	// Returns map[string]int{"a": 2, "b": 4}
func MapValues[K comparable, V1 any, V2 any](m map[K]V1, mapper func(V1) V2) map[K]V2 {
	result := make(map[K]V2, len(m))
	for k, v := range m {
		result[k] = mapper(v)
	}
	return result
}

// Update updates the value at the given key if it exists.
// The updater function receives the current value and returns the new value.
// Returns true if the key was found and updated, false otherwise.
//
// Example:
//
//	m := map[string]int{"counter": 0}
//	Update(m, "counter", func(v int) int { return v + 1 }) // m is now {"counter": 1}
//	Update(m, "missing", func(v int) int { return v + 1 }) // no change, returns false
func Update[K comparable, V any](m map[K]V, key K, updater func(V) V) bool {
	if v, ok := m[key]; ok {
		m[key] = updater(v)
		return true
	}
	return false
}

// GetOrSet retrieves the value for the given key, or if the key doesn't exist,
// computes the value using the provided function, sets it in the map, and returns it.
//
// Example:
//
//	m := map[string]int{}
//	v1 := GetOrSet(m, "key", func() int { return expensiveComputation() })
//	v2 := GetOrSet(m, "key", func() int { return 999 }) // Returns same value as v1, func not called
func GetOrSet[K comparable, V any](m map[K]V, key K, compute func() V) V {
	if v, ok := m[key]; ok {
		return v
	}
	v := compute()
	m[key] = v
	return v
}

// MergeWith merges multiple maps, using the merge function to resolve key conflicts.
// When the same key exists in multiple maps, the merge function is called with all
// values for that key to produce the final value.
//
// Example:
//
//	m1 := map[string]int{"a": 1, "b": 2}
//	m2 := map[string]int{"b": 3, "c": 4}
//	MergeWith(m1, m2, func(v1, v2 int) int { return v1 + v2 })
//	// Returns map[string]int{"a": 1, "b": 5, "c": 4}
func MergeWith[K comparable, V any](maps []map[K]V, merge func(values ...V) V) map[K]V {
	if len(maps) == 0 {
		return map[K]V{}
	}

	// Collect all values for each key
	keyValues := make(map[K][]V)
	capacity := 0
	for _, m := range maps {
		capacity += len(m)
		for k, v := range m {
			keyValues[k] = append(keyValues[k], v)
		}
	}

	// Merge values for each key
	result := make(map[K]V, capacity)
	for k, values := range keyValues {
		result[k] = merge(values...)
	}
	return result
}

// Difference returns a new map containing only the key-value pairs from m1
// that are not present in m2 (keys in m1 but not in m2).
//
// Example:
//
//	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
//	m2 := map[string]int{"b": 20, "d": 4}
//	Difference(m1, m2) // Returns map[string]int{"a": 1, "c": 3}
func Difference[K comparable, V any](m1, m2 map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range m1 {
		if _, exists := m2[k]; !exists {
			result[k] = v
		}
	}
	return result
}

// Intersection returns a new map containing only the key-value pairs
// that exist in both m1 and m2. Values from m1 are used.
//
// Example:
//
//	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
//	m2 := map[string]int{"b": 20, "c": 30, "d": 4}
//	Intersection(m1, m2) // Returns map[string]int{"b": 2, "c": 3}
func Intersection[K comparable, V any](m1, m2 map[K]V) map[K]V {
	result := make(map[K]V)
	// Iterate over the smaller map for efficiency
	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}
	for k, v := range m1 {
		if _, exists := m2[k]; exists {
			result[k] = v
		}
	}
	return result
}
