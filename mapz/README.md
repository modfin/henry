# mapz

> Generic map operations and utilities for Go

The `mapz` package provides 30+ utility functions for working with maps. All functions are generic and immutable (return new maps rather than modifying inputs, unless explicitly marked).

## Quick Reference

**By Category:**
- [Access & Extraction](#access--extraction) - Keys, Values, Lookup
- [Comparison](#comparison) - Equal, EqualBy
- [Cloning & Copying](#cloning--copying) - Clone, Copy, Clear
- [Combination](#combination) - Merge, MergeWith
- [Filtering](#filtering) - Filter, Reject, By Keys/Values
- [Transformation](#transformation) - MapKeys, MapValues, Remap, Invert
- [Update Operations](#update-operations) - Update, GetOrSet
- [Set Operations](#set-operations) - Difference, Intersection
- [Conversion](#conversion) - To/From Slice, Entries

## Installation

```bash
go get github.com/modfin/henry/mapz
```

## Usage

```go
import "github.com/modfin/henry/mapz"

scores := map[string]int{"Alice": 85, "Bob": 92, "Charlie": 78}

// Get all names
names := mapz.Keys(scores)
// names = []string{"Alice", "Bob", "Charlie"} (order not guaranteed)

// Filter high scores
highScorers := mapz.Filter(scores, func(name string, score int) bool {
    return score >= 80
})
// highScorers = {"Alice": 85, "Bob": 92}
```

## Function Categories

### Access & Extraction

#### Keys
Extract all keys as a slice.

```go
scores := map[string]int{"Alice": 85, "Bob": 92}
names := mapz.Keys(scores)
// names = []string{"Alice", "Bob"} (order not guaranteed)
```

#### Values
Extract all values as a slice.

```go
scores := map[string]int{"Alice": 85, "Bob": 92}
scoresList := mapz.Values(scores)
// scoresList = []int{85, 92} (order not guaranteed)
```

#### ValueOr
Get value or fallback.

```go
scores := map[string]int{"Alice": 85}
alice := mapz.ValueOr(scores, "Alice", 0)      // alice = 85
charlie := mapz.ValueOr(scores, "Charlie", 0)  // charlie = 0 (fallback)
```

### Comparison

#### Equal
Check if two maps are equal (keys and values match).

```go
m1 := map[string]int{"a": 1, "b": 2}
m2 := map[string]int{"a": 1, "b": 2}
m3 := map[string]int{"a": 1, "b": 3}

mapz.Equal(m1, m2) // true
mapz.Equal(m1, m3) // false
```

#### EqualBy
Check equality with custom comparison function.

```go
m1 := map[string]int{"a": 1, "b": 2}
m2 := map[string]int{"a": 2, "b": 4}

// Check if values are all even
mapz.EqualBy(m1, m2, func(v1, v2 int) bool {
    return v1%2 == v2%2
}) // true (both 1→2 and 2→4 are odd→even)
```

### Cloning & Copying

#### Clone
Create a shallow copy of a map.

```go
original := map[string]int{"a": 1, "b": 2}
copy := mapz.Clone(original)
// copy = {"a": 1, "b": 2}
// Modifying copy doesn't affect original
```

#### Copy
Copy all entries from source to destination.

```go
dst := map[string]int{"a": 1}
src := map[string]int{"b": 2, "c": 3}
mapz.Copy(dst, src)
// dst = {"a": 1, "b": 2, "c": 3}
```

#### Clear
Remove all entries from a map (mutates).

```go
m := map[string]int{"a": 1, "b": 2}
mapz.Clear(m)
// m = {}
```

### Combination

#### Merge
Merge multiple maps into one.

```go
m1 := map[string]int{"a": 1, "b": 2}
m2 := map[string]int{"b": 20, "c": 3}  // Note: b has different value
m3 := map[string]int{"d": 4}

merged := mapz.Merge(m1, m2, m3)
// merged = {"a": 1, "b": 20, "c": 3, "d": 4}
// Later maps overwrite earlier ones for duplicate keys
```

#### MergeWith
Merge with custom conflict resolution.

```go
m1 := map[string]int{"a": 1, "b": 2}
m2 := map[string]int{"b": 3, "c": 4}

// Sum values for duplicate keys
merged := mapz.MergeWith([]map[string]int{m1, m2}, func(vals ...int) int {
    sum := 0
    for _, v := range vals {
        sum += v
    }
    return sum
})
// merged = {"a": 1, "b": 5, "c": 4}  (2 + 3 = 5)
```

### Filtering

#### Filter
Keep entries matching predicate.

```go
scores := map[string]int{"Alice": 85, "Bob": 72, "Charlie": 90}
highScorers := mapz.Filter(scores, func(name string, score int) bool {
    return score >= 80
})
// highScorers = {"Alice": 85, "Charlie": 90}
```

#### FilterByKeys
Keep only specified keys.

```go
scores := map[string]int{"Alice": 85, "Bob": 72, "Charlie": 90}
selected := mapz.FilterByKeys(scores, []string{"Alice", "Charlie", "Dave"})
// selected = {"Alice": 85, "Charlie": 90}
// Dave is ignored (doesn't exist in original)
```

#### FilterByValues
Keep only entries with specified values.

```go
scores := map[string]int{"Alice": 85, "Bob": 90, "Charlie": 90}
in90s := mapz.FilterByValues(scores, []int{90})
// in90s = {"Bob": 90, "Charlie": 90}
```

#### Reject
Remove entries matching predicate (inverse of Filter).

```go
scores := map[string]int{"Alice": 85, "Bob": 72, "Charlie": 90}
passing := mapz.Reject(scores, func(name string, score int) bool {
    return score < 80
})
// passing = {"Alice": 85, "Charlie": 90}
```

#### RejectByKeys
Remove specified keys.

```go
scores := map[string]int{"Alice": 85, "Bob": 72, "Charlie": 90}
withoutBob := mapz.RejectByKeys(scores, []string{"Bob"})
// withoutBob = {"Alice": 85, "Charlie": 90}
```

#### RejectByValues
Remove entries with specified values.

```go
scores := map[string]int{"Alice": 85, "Bob": 72, "Charlie": 72}
without72s := mapz.RejectByValues(scores, []int{72})
// without72s = {"Alice": 85}
```

#### Delete
Remove entries where predicate is true (mutates).

```go
scores := map[string]int{"Alice": 85, "Bob": 72, "Charlie": 90}
mapz.Delete(scores, func(name string, score int) bool {
    return score < 80
})
// scores = {"Alice": 85, "Charlie": 90}
```

#### DeleteKeys
Remove entries by keys (mutates).

```go
scores := map[string]int{"Alice": 85, "Bob": 72, "Charlie": 90}
mapz.DeleteKeys(scores, "Bob")
// scores = {"Alice": 85, "Charlie": 90}
```

#### DeleteValues
Remove entries by values (mutates).

```go
scores := map[string]int{"Alice": 85, "Bob": 72, "Charlie": 72}
mapz.DeleteValues(scores, 72)
// scores = {"Alice": 85}
```

### Transformation

#### MapKeys
Transform all keys.

```go
scores := map[string]int{"alice": 85, "bob": 92}
upperKeys := mapz.MapKeys(scores, strings.ToUpper)
// upperKeys = {"ALICE": 85, "BOB": 92}
```

#### MapValues
Transform all values.

```go
scores := map[string]int{"Alice": 85, "Bob": 92}
doubled := mapz.MapValues(scores, func(score int) int {
    return score * 2
})
// doubled = {"Alice": 170, "Bob": 184}
```

#### Remap
Transform both keys and values.

```go
scores := map[string]int{"Alice": 85, "Bob": 92}
remapped := mapz.Remap(scores, func(name string, score int) (string, string) {
    return strings.ToUpper(name), fmt.Sprintf("%d%%", score)
})
// remapped = {"ALICE": "85%", "BOB": "92%"}
```

#### RemapKeys
Transform only keys.

```go
ids := map[int]string{1: "Alice", 2: "Bob"}
withPrefix := mapz.RemapKeys(ids, func(id int, name string) string {
    return fmt.Sprintf("user-%d", id)
})
// withPrefix = {"user-1": "Alice", "user-2": "Bob"}
```

#### RemapValues
Transform only values (like MapValues but receives key too).

```go
scores := map[string]int{"Alice": 85, "Bob": 92}
withGrades := mapz.RemapValues(scores, func(name string, score int) string {
    if score >= 90 {
        return "A"
    } else if score >= 80 {
        return "B"
    }
    return "C"
})
// withGrades = {"Alice": "B", "Bob": "A"}
```

#### Invert
Swap keys and values.

```go
scores := map[string]int{"Alice": 85, "Bob": 92}
byScore := mapz.Invert(scores)
// byScore = map[int]string{85: "Alice", 92: "Bob"}
// Note: If duplicate values exist, last one wins
```

### Update Operations

#### Update
Update value at key if it exists.

```go
counters := map[string]int{"visits": 0, "clicks": 5}
updated := mapz.Update(counters, "visits", func(v int) int {
    return v + 1
})
// updated = true, counters = {"visits": 1, "clicks": 5}

// Key doesn't exist
updated = mapz.Update(counters, "downloads", func(v int) int {
    return v + 1
})
// updated = false, counters unchanged
```

#### GetOrSet
Get value or compute and set if missing.

```go
cache := map[string]string{}

// First call - computes and caches
value := mapz.GetOrSet(cache, "config", func() string {
    // Expensive operation
    return loadConfigFromDisk()
})

// Second call - returns cached value
value = mapz.GetOrSet(cache, "config", func() string {
    // This function is NOT called!
    return "never used"
})
```

### Set Operations

#### Difference
Keys in first map but not in second.

```go
m1 := map[string]int{"a": 1, "b": 2, "c": 3}
m2 := map[string]int{"b": 20, "d": 4}
diff := mapz.Difference(m1, m2)
// diff = {"a": 1, "c": 3}
```

#### Intersection
Keys present in both maps.

```go
m1 := map[string]int{"a": 1, "b": 2, "c": 3}
m2 := map[string]int{"b": 20, "c": 30, "d": 4}
common := mapz.Intersection(m1, m2)
// common = {"b": 2, "c": 3}
// Values from m1 are used
```

### Conversion

#### Slice
Convert map to slice.

```go
scores := map[string]int{"Alice": 85, "Bob": 92}
pairs := mapz.Slice(scores, func(name string, score int) string {
    return fmt.Sprintf("%s=%d", name, score)
})
// pairs = []string{"Alice=85", "Bob=92"} (order not guaranteed)
```

#### Entry
Type representing a key-value pair.

```go
type Entry[K comparable, V any] struct {
    Key   K
    Value V
}
```

#### Entries
Convert map to slice of entries.

```go
scores := map[string]int{"Alice": 85, "Bob": 92}
entries := mapz.Entries(scores)
// entries = []Entry{{"Alice", 85}, {"Bob", 92}}
```

#### FromEntries
Convert entries back to map.

```go
entries := []mapz.Entry[string, int]{
    {Key: "Alice", Value: 85},
    {Key: "Bob", Value: 92},
}
m := mapz.FromEntries(entries)
// m = {"Alice": 85, "Bob": 92}
```

## Performance Notes

- **Immutable by default**: Functions return new maps (safer, easier to reason about)
- **Pre-allocation**: Merge and Clone pre-allocate capacity to avoid reallocations
- **Mutation marked**: Functions that mutate (Clear, Delete, etc.) are clearly documented

## See Also

- [setz](../setz/) - Dedicated set data structure
- [slicez](../slicez/) - Work with slice representations of map data
