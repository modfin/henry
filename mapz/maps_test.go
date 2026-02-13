package mapz

import (
	"fmt"
	"github.com/modfin/henry/slicez"
	"reflect"
	"strings"
	"testing"
)

func TestEqual(t *testing.T) {
	m1 := map[int]int{1: 1, 3: 3}
	m2 := map[int]int{1: 1, 3: 3}
	if !Equal(m1, m2) {
		t.Log("Expected maps to be equal")
		t.Fail()
	}

	m2[4] = 4
	if Equal(m1, m2) {
		t.Log("Expected maps not to be equal")
		t.Fail()
	}

	m1[4] = 3
	if Equal(m1, m2) {
		t.Log("Expected maps not to be equal")
		t.Fail()
	}
}

func TestClear(t *testing.T) {
	m1 := map[int]int{1: 1, 3: 3}
	m2 := map[int]int{}
	Clear(m1)
	if !Equal(m1, m2) {
		t.Log("Expected maps be equal")
		t.Fail()
	}
}

func TestClone(t *testing.T) {
	m1 := map[int]int{1: 1, 3: 3}
	m2 := Clone(m1)
	if &m1 == &m2 {
		t.Log("should not be the same")
		t.Fail()
	}
	if !Equal(m1, m2) {
		t.Log("Expected maps be equal")
		t.Fail()
	}
}

func TestCopy(t *testing.T) {
	src := map[int]int{1: 1, 3: 3}
	dst := map[int]int{3: 33, 4: 4}
	exp := map[int]int{1: 1, 3: 3, 4: 4}
	Copy(dst, src)
	if !Equal(exp, dst) {
		t.Log("Expected maps be equal")
		t.Fail()
	}
}

func TestValues(t *testing.T) {
	m := map[int]int{1: 11, 3: 33}
	v := Values(m)

	if len(v) != 2 {
		t.Log("Expected size of 2")
		t.Fail()
	}

	if !slicez.Contains(v, 33) {
		t.Log("Expected value 11")
		t.Fail()
	}
	if !slicez.Contains(v, 11) {
		t.Log("Expected value 31")
		t.Fail()
	}
}

func TestMerge(t *testing.T) {
	m1 := map[int]int{1: 11, 3: 33}
	m2 := map[int]int{2: 22, 4: 44}
	exp := map[int]int{1: 11, 3: 33, 2: 22, 4: 44}
	m := Merge(m1, m2)

	if !Equal(m, exp) {
		t.Log("Expected equality")
		t.Fail()
	}

	m1[4] = 444
	m = Merge(m1, m2)
	if !Equal(m, exp) {
		t.Log("Expected equality")
		t.Fail()
	}

	exp[4] = 444
	m = Merge(m2, m1)
	if !Equal(m, exp) {
		t.Log("Expected equality")
		t.Fail()
	}

}

func TestRemap(t *testing.T) {
	m := map[int]int{1: 11, 3: 33}
	exp := map[string]int{"1": 110, "3": 330}

	res := Remap(m, func(k int, v int) (string, int) {
		return fmt.Sprint(k), v * 10
	})

	if !Equal(res, exp) {
		t.Log("Expected equality")
		t.Fail()
	}
}

func TestKeys(t *testing.T) {
	m := map[int]int{1: 11, 3: 33}
	v := Keys(m)

	if len(v) != 2 {
		t.Log("Expected size of 2")
		t.Fail()
	}

	if !slicez.Contains(v, 1) {
		t.Log("Expected value 1")
		t.Fail()
	}
	if !slicez.Contains(v, 3) {
		t.Log("Expected value 3")
		t.Fail()
	}
}

func TestDeleteValue(t *testing.T) {
	m := map[int]int{1: 11, 3: 33}
	DeleteValues(m, 33)

	if len(m) != 1 {
		t.Log("Expected size of 1")
		t.Fail()
	}

	if m[1] != 11 {
		t.Log("Expected to contain key 1")
		t.Fail()
	}
}

func TestDeleteFunc(t *testing.T) {
	m := map[int]int{1: 11, 3: 33}
	Delete(m, func(k int, v int) bool {
		return k == 3
	})

	if len(m) != 1 {
		t.Log("Expected size of 1")
		t.Fail()
	}

	if m[1] != 11 {
		t.Log("Expected to contain key 1")
		t.Fail()
	}
}

func TestPickByKeys(t *testing.T) {
	m := FilterByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz"})
	exp := map[string]int{"foo": 1, "baz": 3}
	if !reflect.DeepEqual(m, exp) {
		t.Log("Expected", exp, "got", m)
		t.Fail()
	}
}
func TestPickByValues(t *testing.T) {
	m := FilterByValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []int{23, 2})
	exp := map[string]int{"bar": 2}
	if !reflect.DeepEqual(m, exp) {
		t.Log("Expected", exp, "got", m)
		t.Fail()
	}
}
func TestOmitByKeys(t *testing.T) {
	m := RejectByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz"})
	exp := map[string]int{"bar": 2}
	if !reflect.DeepEqual(m, exp) {
		t.Log("Expected", exp, "got", m)
		t.Fail()
	}
}
func TestOmitByValues(t *testing.T) {
	m := RejectByValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []int{23, 2})
	exp := map[string]int{"foo": 1, "baz": 3}
	if !reflect.DeepEqual(m, exp) {
		t.Log("Expected", exp, "got", m)
		t.Fail()
	}
}

func TestEqualBy(t *testing.T) {
	m1 := map[int]string{1: "one", 2: "two"}
	m2 := map[int]string{1: "ONE", 2: "TWO"}
	// Case-insensitive comparison
	eq := func(a, b string) bool {
		return strings.ToLower(a) == strings.ToLower(b)
	}
	if !EqualBy(m1, m2, eq) {
		t.Error("Expected EqualBy to return true for case-insensitive match")
	}
	// Different values
	m3 := map[int]string{1: "one", 2: "different"}
	if EqualBy(m1, m3, eq) {
		t.Error("Expected EqualBy to return false for different values")
	}
}

func TestDeleteKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	DeleteKeys(m, "b", "d")
	if len(m) != 2 {
		t.Errorf("Expected length 2, got %d", len(m))
	}
	if _, ok := m["a"]; !ok {
		t.Error("Expected key 'a' to exist")
	}
	if _, ok := m["c"]; !ok {
		t.Error("Expected key 'c' to exist")
	}
}

func TestFilter(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	result := Filter(m, func(k string, v int) bool {
		return v%2 == 0
	})
	if len(result) != 2 {
		t.Errorf("Expected length 2, got %d", len(result))
	}
	if result["b"] != 2 || result["d"] != 4 {
		t.Error("Expected only even values")
	}
}

func TestReject(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	result := Reject(m, func(k string, v int) bool {
		return v%2 == 0
	})
	if len(result) != 2 {
		t.Errorf("Expected length 2, got %d", len(result))
	}
	if result["a"] != 1 || result["c"] != 3 {
		t.Error("Expected only odd values")
	}
}

func TestSlice(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	result := Slice(m, func(k string, v int) string {
		return fmt.Sprintf("%s=%d", k, v)
	})
	if len(result) != 2 {
		t.Errorf("Expected length 2, got %d", len(result))
	}
	// Check that all expected values are present
	found := make(map[string]bool)
	for _, s := range result {
		found[s] = true
	}
	if !found["a=1"] || !found["b=2"] {
		t.Errorf("Expected [a=1, b=2], got %v", result)
	}
}

func TestEntries(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	entries := Entries(m)
	if len(entries) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(entries))
	}
	// Convert to map for easier checking
	result := make(map[string]int)
	for _, e := range entries {
		result[e.Key] = e.Value
	}
	if !reflect.DeepEqual(result, m) {
		t.Errorf("Entries() = %v, want %v", result, m)
	}
}

func TestFromEntries(t *testing.T) {
	entries := []Entry[string, int]{
		{Key: "a", Value: 1},
		{Key: "b", Value: 2},
	}
	result := FromEntries(entries)
	expected := map[string]int{"a": 1, "b": 2}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FromEntries() = %v, want %v", result, expected)
	}
}

func TestRemapKeys(t *testing.T) {
	m := map[int]int{1: 10, 2: 20}
	result := RemapKeys(m, func(k, v int) string {
		return fmt.Sprintf("key_%d", k)
	})
	expected := map[string]int{"key_1": 10, "key_2": 20}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("RemapKeys() = %v, want %v", result, expected)
	}
}

func TestRemapValues(t *testing.T) {
	m := map[int]int{1: 10, 2: 20}
	result := RemapValues(m, func(k, v int) string {
		return fmt.Sprintf("value_%d", v)
	})
	expected := map[int]string{1: "value_10", 2: "value_20"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("RemapValues() = %v, want %v", result, expected)
	}
}

func TestInvert(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	result := Invert(m)
	expected := map[int]string{1: "a", 2: "b", 3: "c"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Invert() = %v, want %v", result, expected)
	}
}

func TestValueOr(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	// Key exists
	if ValueOr(m, "a", 100) != 1 {
		t.Error("Expected ValueOr to return 1 for key 'a'")
	}
	// Key does not exist
	if ValueOr(m, "c", 100) != 100 {
		t.Error("Expected ValueOr to return fallback 100 for key 'c'")
	}
}

func TestMapKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	result := MapKeys(m, strings.ToUpper)
	expected := map[string]int{"A": 1, "B": 2}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MapKeys() = %v, want %v", result, expected)
	}

	// Empty map
	empty := MapKeys(map[string]int{}, strings.ToUpper)
	if len(empty) != 0 {
		t.Error("MapKeys on empty map should return empty map")
	}
}

func TestMapValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	result := MapValues(m, func(v int) int { return v * 2 })
	expected := map[string]int{"a": 2, "b": 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MapValues() = %v, want %v", result, expected)
	}

	// Empty map
	empty := MapValues(map[string]int{}, func(v int) int { return v * 2 })
	if len(empty) != 0 {
		t.Error("MapValues on empty map should return empty map")
	}
}

func TestUpdate(t *testing.T) {
	m := map[string]int{"counter": 0, "value": 10}

	// Update existing key
	updated := Update(m, "counter", func(v int) int { return v + 1 })
	if !updated {
		t.Error("Update should return true for existing key")
	}
	if m["counter"] != 1 {
		t.Errorf("After Update, m[\"counter\"] = %d, want 1", m["counter"])
	}

	// Try to update non-existing key
	notUpdated := Update(m, "missing", func(v int) int { return v + 1 })
	if notUpdated {
		t.Error("Update should return false for non-existing key")
	}
	if _, exists := m["missing"]; exists {
		t.Error("Update should not add missing key")
	}
}

func TestGetOrSet(t *testing.T) {
	m := map[string]int{}
	callCount := 0
	compute := func() int {
		callCount++
		return 42
	}

	// First call - should compute
	v1 := GetOrSet(m, "key", compute)
	if v1 != 42 {
		t.Errorf("GetOrSet first call = %d, want 42", v1)
	}
	if callCount != 1 {
		t.Errorf("Compute function called %d times, expected 1", callCount)
	}

	// Second call - should return cached value
	v2 := GetOrSet(m, "key", compute)
	if v2 != 42 {
		t.Errorf("GetOrSet second call = %d, want 42", v2)
	}
	if callCount != 1 {
		t.Errorf("Compute function called %d times, expected 1 (cached)", callCount)
	}

	// Verify key was set
	if m["key"] != 42 {
		t.Errorf("m[\"key\"] = %d, want 42", m["key"])
	}
}

func TestMergeWith(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}

	// Sum values for duplicate keys
	result := MergeWith([]map[string]int{m1, m2}, func(values ...int) int {
		sum := 0
		for _, v := range values {
			sum += v
		}
		return sum
	})

	expected := map[string]int{"a": 1, "b": 5, "c": 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MergeWith() = %v, want %v", result, expected)
	}

	// Empty maps slice
	empty := MergeWith([]map[string]int{}, func(values ...int) int { return 0 })
	if len(empty) != 0 {
		t.Error("MergeWith with empty slice should return empty map")
	}

	// Single map
	single := MergeWith([]map[string]int{{"x": 10}}, func(values ...int) int { return values[0] })
	if !reflect.DeepEqual(single, map[string]int{"x": 10}) {
		t.Error("MergeWith with single map should return that map")
	}
}

func TestDifference(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"b": 20, "d": 4}

	result := Difference(m1, m2)
	expected := map[string]int{"a": 1, "c": 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Difference() = %v, want %v", result, expected)
	}

	// Empty maps
	if len(Difference(map[string]int{}, m2)) != 0 {
		t.Error("Difference with empty first map should return empty map")
	}
	if !reflect.DeepEqual(Difference(m1, map[string]int{}), m1) {
		t.Error("Difference with empty second map should return first map")
	}
}

func TestIntersection(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"b": 20, "c": 30, "d": 4}

	result := Intersection(m1, m2)
	// Values from m1 should be used
	expected := map[string]int{"b": 2, "c": 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Intersection() = %v, want %v", result, expected)
	}

	// Empty maps
	if len(Intersection(map[string]int{}, m2)) != 0 {
		t.Error("Intersection with empty first map should return empty map")
	}
	if len(Intersection(m1, map[string]int{})) != 0 {
		t.Error("Intersection with empty second map should return empty map")
	}

	// No intersection
	noIntersect := Intersection(map[string]int{"a": 1}, map[string]int{"b": 2})
	if len(noIntersect) != 0 {
		t.Error("Intersection with no common keys should return empty map")
	}
}
