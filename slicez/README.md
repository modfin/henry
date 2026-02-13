# slicez

> Generic slice operations and algorithms for Go

The `slicez` package provides 70+ utility functions for working with slices. All functions are generic and work with any type.

## Quick Reference

**By Category:**
- [Transformation](#transformation) - Map, FlatMap, Reverse, Shuffle
- [Filtering](#filtering) - Filter, Reject, Take, Drop, Compact
- [Searching](#searching) - Contains, Find, Index, Search
- [Aggregation](#aggregation) - Fold, Reduce, Every, Some, None
- [Set Operations](#set-operations) - Union, Intersection, Difference, Uniq
- [Sorting](#sorting) - Sort, SortBy, IsSorted, Max, Min
- [Grouping](#grouping) - GroupBy, Partition, Chunk, ChunkBy
- [Combining](#combining) - Zip, Unzip, Interleave, Concat
- [Utilities](#utilities) - Clone, Sample, Fill, Range, Repeat

## Installation

```bash
go get github.com/modfin/henry/slicez
```

## Usage

```go
import "github.com/modfin/henry/slicez"

numbers := []int{1, 2, 3, 4, 5}
doubled := slicez.Map(numbers, func(n int) int { return n * 2 })
// doubled = []int{2, 4, 6, 8, 10}
```

## Function Categories

### Transformation

Transform elements from one form to another.

#### Map
Transform each element.

```go
// Convert integers to strings
nums := []int{1, 2, 3}
strings := slicez.Map(nums, func(n int) string {
    return fmt.Sprintf("num-%d", n)
})
// strings = []string{"num-1", "num-2", "num-3"}
```

#### FlatMap
Map then flatten results.

```go
// Split each string into words
lines := []string{"hello world", "foo bar"}
words := slicez.FlatMap(lines, func(s string) []string {
    return strings.Split(s, " ")
})
// words = []string{"hello", "world", "foo", "bar"}
```

#### Reverse
Reverse element order.

```go
nums := []int{1, 2, 3, 4, 5}
reversed := slicez.Reverse(nums)
// reversed = []int{5, 4, 3, 2, 1}
```

#### Shuffle
Randomize element order (Fisher-Yates algorithm).

```go
cards := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
shuffled := slicez.Shuffle(cards)
// shuffled = []int{7, 2, 9, 1, 5, 3, 8, 4, 10, 6} (random order)
```

### Filtering

Select or remove elements based on criteria.

#### Filter
Keep elements matching predicate.

```go
nums := []int{1, 2, 3, 4, 5, 6}
evens := slicez.Filter(nums, func(n int) bool {
    return n%2 == 0
})
// evens = []int{2, 4, 6}
```

#### Reject
Remove elements matching predicate (inverse of Filter).

```go
nums := []int{1, 2, 3, 4, 5, 6}
odds := slicez.Reject(nums, func(n int) bool {
    return n%2 == 0
})
// odds = []int{1, 3, 5}
```

#### Take
Take first N elements.

```go
nums := []int{1, 2, 3, 4, 5}
first3 := slicez.Take(nums, 3)
// first3 = []int{1, 2, 3}
```

#### TakeWhile
Take elements while predicate is true.

```go
nums := []int{1, 2, 3, 4, 5, 1, 2}
ascending := slicez.TakeWhile(nums, func(n int) bool {
    return n < 4
})
// ascending = []int{1, 2, 3}
```

#### Drop
Drop first N elements.

```go
nums := []int{1, 2, 3, 4, 5}
rest := slicez.Drop(nums, 2)
// rest = []int{3, 4, 5}
```

#### DropWhile
Drop elements while predicate is true.

```go
nums := []int{1, 2, 3, 4, 5}
from4 := slicez.DropWhile(nums, func(n int) bool {
    return n < 4
})
// from4 = []int{4, 5}
```

#### Compact
Remove consecutive duplicates.

```go
data := []int{1, 1, 2, 2, 2, 3, 3}
compacted := slicez.Compact(data)
// compacted = []int{1, 2, 3}
```

#### Deduplicate
Remove all consecutive duplicates (alias for Compact).

```go
data := []int{1, 1, 2, 1, 2, 2, 3}
result := slicez.Deduplicate(data)
// result = []int{1, 2, 1, 2, 3}
```

### Searching

Find elements and their positions.

#### Contains
Check if element exists (for comparable types).

```go
nums := []int{1, 2, 3, 4, 5}
has3 := slicez.Contains(nums, 3)     // true
has7 := slicez.Contains(nums, 7)     // false
```

#### ContainsBy
Check if element exists using predicate.

```go
users := []User{{Name: "Alice", Age: 30}, {Name: "Bob", Age: 25}}
hasAdult := slicez.ContainsBy(users, func(u User) bool {
    return u.Age >= 18
}) // true
```

#### Find
Find first element matching predicate.

```go
nums := []int{1, 2, 3, 4, 5, 6, 7}
firstEven, found := slicez.Find(nums, func(n int) bool {
    return n%2 == 0
})
// firstEven = 2, found = true
```

#### FindLast
Find last element matching predicate.

```go
nums := []int{1, 2, 3, 4, 5, 6, 7}
lastEven, found := slicez.FindLast(nums, func(n int) bool {
    return n%2 == 0
})
// lastEven = 6, found = true
```

#### Index
Find index of first occurrence (for comparable types).

```go
nums := []int{10, 20, 30, 40, 50}
idx := slicez.Index(nums, 30)   // idx = 2
idx = slicez.Index(nums, 100)   // idx = -1 (not found)
```

#### IndexBy
Find index using predicate.

```go
nums := []int{1, 2, 3, 4, 5}
idx := slicez.IndexBy(nums, func(n int) bool {
    return n > 3
}) // idx = 3 (first element > 3)
```

#### Search
Binary search on sorted slice.

```go
nums := []int{1, 3, 5, 7, 9, 11, 13}
idx, val := slicez.Search(nums, func(n int) bool {
    return n >= 7
})
// idx = 3, val = 7
```

### Aggregation

Reduce slices to single values or check properties.

#### Fold / Reduce
Reduce from left to single value.

```go
nums := []int{1, 2, 3, 4, 5}
sum := slicez.Fold(nums, func(acc, val int) int {
    return acc + val
}, 0)
// sum = 15
```

#### FoldRight
Reduce from right to single value.

```go
words := []string{"a", "b", "c"}
result := slicez.FoldRight(words, func(acc, val string) string {
    if acc == "" { return val }
    return val + "-" + acc
}, "")
// result = "a-b-c"
```

#### Every
Check if all elements equal value (comparable types).

```go
nums := []int{5, 5, 5}
all5 := slicez.Every(nums, 5)   // true
```

#### EveryBy
Check if all elements satisfy predicate.

```go
nums := []int{2, 4, 6, 8, 10}
allEven := slicez.EveryBy(nums, func(n int) bool {
    return n%2 == 0
}) // true
```

#### Some
Check if any element equals value (alias for Contains).

```go
nums := []int{1, 2, 3, 4, 5}
hasEven := slicez.Some(nums, 2)   // true
```

#### SomeBy
Check if any element satisfies predicate.

```go
nums := []int{1, 3, 5, 6, 7}
hasEven := slicez.SomeBy(nums, func(n int) bool {
    return n%2 == 0
}) // true (6 is even)
```

#### None
Check if no element equals value.

```go
nums := []int{1, 2, 3}
none0 := slicez.None(nums, 0)   // true
```

#### NoneBy
Check if no element satisfies predicate.

```go
nums := []int{1, 2, 3}
noNegatives := slicez.NoneBy(nums, func(n int) bool {
    return n < 0
}) // true
```

### Set Operations

Work with sets (unique collections).

#### Uniq
Remove all duplicates.

```go
nums := []int{1, 2, 2, 3, 3, 3, 4}
unique := slicez.Uniq(nums)
// unique = []int{1, 2, 3, 4}
```

#### UniqBy
Remove duplicates by key function.

```go
type Person struct { Name string; City string }
people := []Person{
    {"Alice", "NYC"}, {"Bob", "LA"}, {"Charlie", "NYC"},
}
// Uniq by city (keeps first occurrence)
byCity := slicez.UniqBy(people, func(p Person) string {
    return p.City
})
// byCity = [{"Alice", "NYC"}, {"Bob", "LA"}]
```

#### Union
Combine multiple slices, removing duplicates.

```go
a := []int{1, 2, 3}
b := []int{2, 3, 4}
c := []int{3, 4, 5}
combined := slicez.Union(a, b, c)
// combined = []int{1, 2, 3, 4, 5}
```

#### Intersection
Keep only elements present in ALL slices.

```go
a := []int{1, 2, 3, 4}
b := []int{2, 3, 4, 5}
c := []int{3, 4, 5, 6}
common := slicez.Intersection(a, b, c)
// common = []int{3, 4}
```

#### Difference
Keep only elements NOT present in all slices.

```go
a := []int{1, 2, 3, 4}
b := []int{2, 3}
diff := slicez.Difference(a, b)
// diff = []int{1, 4}
```

#### Complement
Elements in second slice not in first.

```go
seen := []int{1, 2, 3}
all := []int{1, 2, 3, 4, 5, 6}
newOnes := slicez.Complement(seen, all)
// newOnes = []int{4, 5, 6}
```

#### XOR
Symmetric difference (elements in exactly one slice).

```go
a := []int{1, 2, 3}
b := []int{2, 3, 4}
xor := slicez.XOR(a, b)
// xor = []int{1, 4}
```

### Sorting

Sort and find extrema.

#### Sort
Sort using natural order.

```go
nums := []int{3, 1, 4, 1, 5, 9, 2, 6}
sorted := slicez.Sort(nums)
// sorted = []int{1, 1, 2, 3, 4, 5, 6, 9}
```

#### SortBy
Sort using custom comparison.

```go
type Person struct { Name string; Age int }
people := []Person{
    {"Alice", 30}, {"Bob", 25}, {"Charlie", 35},
}
byAge := slicez.SortBy(people, func(a, b Person) bool {
    return a.Age < b.Age
})
// byAge = [{Bob 25} {Alice 30} {Charlie 35}]
```

#### IsSorted
Check if slice is sorted.

```go
nums := []int{1, 2, 3, 4, 5}
slicez.IsSorted(nums)   // true
```

#### Max
Find maximum value.

```go
max := slicez.Max(3, 1, 4, 1, 5, 9, 2, 6)
// max = 9
```

#### Min
Find minimum value.

```go
min := slicez.Min(3, 1, 4, 1, 5, 9, 2, 6)
// min = 1
```

### Grouping

Organize elements into groups.

#### GroupBy
Group by key function.

```go
type Person struct { Name string; Age int }
people := []Person{
    {"Alice", 30}, {"Bob", 25}, {"Charlie", 30},
}
byAge := slicez.GroupBy(people, func(p Person) int {
    return p.Age
})
// byAge = map[int][]Person{
//     30: {{"Alice", 30}, {"Charlie", 30}},
//     25: {{"Bob", 25}},
// }
```

#### Partition
Split into two groups by predicate.

```go
nums := []int{1, 2, 3, 4, 5, 6}
evens, odds := slicez.Partition(nums, func(n int) bool {
    return n%2 == 0
})
// evens = []int{2, 4, 6}
// odds = []int{1, 3, 5}
```

#### Chunk
Split into chunks of size N.

```go
nums := []int{1, 2, 3, 4, 5, 6, 7}
chunks := slicez.Chunk(nums, 3)
// chunks = [][]int{{1, 2, 3}, {4, 5, 6}, {7}}
```

#### ChunkBy
Group consecutive elements satisfying predicate.

```go
nums := []int{1, 2, 3, 2, 2, 1}
groups := slicez.ChunkBy(nums, func(a, b int) bool {
    return a <= b  // Group while ascending
})
// groups = [][]int{{1, 2, 3}, {2, 2}, {1}}
```

### Combining

Combine multiple slices.

#### Concat
Concatenate slices.

```go
a := []int{1, 2, 3}
b := []int{4, 5, 6}
c := []int{7, 8, 9}
combined := slicez.Concat(a, b, c)
// combined = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
```

#### Zip
Combine two slices element-wise.

```go
nums := []int{1, 2, 3}
strs := []string{"a", "b", "c"}
zipped := slicez.Zip(nums, strs, func(n int, s string) string {
    return fmt.Sprintf("%d:%s", n, s)
})
// zipped = []string{"1:a", "2:b", "3:c"}
```

#### Zip2
Combine three slices.

```go
a := []int{1, 2, 3}
b := []int{10, 20, 30}
c := []string{"x", "y", "z"}
result := slicez.Zip2(a, b, c, func(x, y int, z string) string {
    return fmt.Sprintf("%d-%d-%s", x, y, z)
})
// result = []string{"1-10-x", "2-20-y", "3-30-z"}
```

#### Unzip
Split slice of pairs into two slices.

```go	type Pair struct{ A, B int }
pairs := []Pair{{1, 10}, {2, 20}, {3, 30}}
as, bs := slicez.Unzip(pairs, func(p Pair) (int, int) {
    return p.A, p.B
})
// as = []int{1, 2, 3}
// bs = []int{10, 20, 30}
```

#### Interleave
Interleave multiple slices round-robin.

```go
a := []int{1, 4, 7}
b := []int{2, 5, 8}
c := []int{3, 6, 9}
interleaved := slicez.Interleave(a, b, c)
// interleaved = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
```

### Utilities

#### Clone
Create a copy of a slice.

```go
original := []int{1, 2, 3}
copy := slicez.Clone(original)
// copy = []int{1, 2, 3}
```

#### Sample
Get N random elements.

```go
nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
sample := slicez.Sample(nums, 3)
// sample = []int{7, 2, 9} (3 random elements)
```

#### Fill
Create slice filled with value.

```go
fives := slicez.Fill(5, 5)
// fives = []int{5, 5, 5, 5, 5}
```

#### Range
Create range of integers.

```go
nums := slicez.Range(1, 5)
// nums = []int{1, 2, 3, 4, 5}
```

#### RangeStep
Create range with step.

```go
evens := slicez.RangeStep(0, 10, 2)
// evens = []int{0, 2, 4, 6, 8, 10}
```

#### Repeat
Repeat slice N times.

```go
pattern := []int{1, 2}
repeated := slicez.Repeat(pattern, 3)
// repeated = []int{1, 2, 1, 2, 1, 2}
```

#### Head / Tail / Initial / Last
Access slice ends.

```go
nums := []int{1, 2, 3, 4, 5}

head, _ := slicez.Head(nums)      // head = 1
tail := slicez.Tail(nums)         // tail = []int{2, 3, 4, 5}
initial := slicez.Initial(nums)   // initial = []int{1, 2, 3, 4}
last, _ := slicez.Last(nums)      // last = 5
```

#### Nth
Get element at index (wraps around).

```go
nums := []int{10, 20, 30, 40, 50}
slicez.Nth(nums, 0)   // 10 (first)
slicez.Nth(nums, 2)   // 30
slicez.Nth(nums, -1)  // 50 (last)
slicez.Nth(nums, 5)   // 10 (wraps around)
```

#### KeyBy
Create map from slice.

```go	type User struct { ID int; Name string }
users := []User{{1, "Alice"}, {2, "Bob"}}
byID := slicez.KeyBy(users, func(u User) int {
    return u.ID
})
// byID = map[int]User{1: {1, "Alice"}, 2: {2, "Bob"}}
```

#### Associate
Convert slice to map (key-value pairs).

```go
nums := []int{1, 2, 3, 4, 5}
squares := slicez.Associate(nums, func(n int) (int, int) {
    return n, n * n
})
// squares = map[int]int{1: 1, 2: 4, 3: 9, 4: 16, 5: 25}
```

#### SliceToMap
Alias for Associate.

```go
// Same as Associate
result := slicez.SliceToMap(nums, mapper)
```

#### Flatten
Flatten 2D slice to 1D.

```go	nested := [][]int{{1, 2}, {3, 4}, {5, 6}}
flat := slicez.Flatten(nested)
// flat = []int{1, 2, 3, 4, 5, 6}
```

#### Join
Join 2D slice with separator.

```go	nested := [][]string{{"a", "b"}, {"c", "d"}}
joined := slicez.Join(nested, []string{"-"})
// joined = []string{"a", "b", "-", "c", "d"}
```

#### Intersperse
Insert element between elements.

```go
nums := []int{1, 2, 3}
withCommas := slicez.Intersperse(nums, 0)
// withCommas = []int{1, 0, 2, 0, 3}
```

#### Transpose
Transpose matrix (rows become columns).

```go
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
}
transposed := slicez.Transpose(matrix)
// transposed = [][]int{
//     {1, 4},
//     {2, 5},
//     {3, 6},
// }
```

#### SlidingWindow
Create sliding windows.

```go
nums := []int{1, 2, 3, 4, 5}
windows := slicez.SlidingWindow(nums, 3)
// windows = [][]int{
//     {1, 2, 3},
//     {2, 3, 4},
//     {3, 4, 5},
// }
```

#### SplitAt
Split at index.

```go
nums := []int{1, 2, 3, 4, 5}
left, right := slicez.SplitAt(nums, 2)
// left = []int{1, 2}
// right = []int{3, 4, 5}
```

#### Span
Split where predicate becomes false.

```go
nums := []int{1, 2, 3, 4, 5}
init, rest := slicez.Span(nums, func(n int) bool {
    return n < 4
})
// init = []int{1, 2, 3}
// rest = []int{4, 5}
```

#### ScanLeft / ScanRight
Running totals.

```go
nums := []int{1, 2, 3, 4}
running := slicez.ScanLeft(nums, func(acc, n int) int {
    return acc + n
}, 0)
// running = []int{0, 1, 3, 6, 10}
```

#### IsAllUnique
Check if all elements are unique.

```go
slicez.IsAllUnique([]int{1, 2, 3})    // true
slicez.IsAllUnique([]int{1, 2, 2})    // false
```

#### ForEach / ForEachRight
Apply function to each element.

```go
slicez.ForEach([]int{1, 2, 3}, func(n int) {
    fmt.Println(n)
})
// Prints 1, 2, 3
```

#### Cut
Split at first occurrence.

```go
nums := []int{1, 2, 3, 4, 5}
left, right, found := slicez.Cut(nums, 3)
// left = []int{1, 2}
// right = []int{4, 5}
// found = true
```

#### Replace
Replace elements.

```go
nums := []int{1, 2, 3, 2, 4}
replaced := slicez.Replace(nums, 2, 9, 1) // Replace 1 occurrence of 2 with 9
// replaced = []int{1, 9, 3, 2, 4}
```

#### Without
Remove specific values.

```go
nums := []int{1, 2, 3, 2, 4, 2}
result := slicez.Without(nums, 2)
// result = []int{1, 3, 4}
```

## Performance Notes

- **Pre-allocation**: Functions like `Map`, `Filter`, `Union` pre-allocate result slices
- **Efficient sampling**: `Sample` uses Fisher-Yates or swap-to-end algorithms
- **In-place mutations**: Some functions (marked in docs) modify input for performance
- **Zero-allocation**: Type aliases (Pipe, Set) have no overhead

## See Also

- [pipez](../pipez/) - Method-chaining API for slice operations
- [compare](../compare/) - Comparison utilities and constraints
