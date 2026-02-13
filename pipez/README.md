# pipez

> Fluent, chainable API for slice operations

The `pipez` package wraps the `slicez` functions in a method-chaining API. It's implemented as a type alias (zero overhead) that provides a more readable, pipeline-style interface.

## Quick Reference

All methods chain and return `Pipe[A]` except terminal operations which return concrete values:

**Chain Methods** (return `Pipe[A]`):
- Transform: `Map`, `Reverse`, `Shuffle`
- Filter: `Filter`, `Reject`, `Take`, `Drop`
- Access: `Tail`, `Initial`
- Combine: `Concat`, `Zip`, `Interleave`

**Terminal Operations** (return values):
- `Slice()` - Get underlying `[]A`
- `Head()` - Get first element `(A, error)`
- `Last()` - Get last element `(A, error)`
- `Count()` - Get length `int`
- `Nth()` - Get element at index `A`
- `Every()`, `Some()`, `None()` - Predicates `bool`
- `Fold()`, `FoldRight()` - Reduce to value
- `Partition()` - Split into two `([]A, []A)`

## Installation

```bash
go get github.com/modfin/henry/pipez
```

## Usage

```go
import "github.com/modfin/henry/pipez"

// Traditional approach
result := slicez.Take(
    slicez.Filter(
        slicez.Map(nums, func(n int) int { return n * 2 }),
        func(n int) bool { return n > 10 },
    ),
    3,
)

// Fluent approach
result := pipez.Of(nums).
    Map(func(n int) int { return n * n }).
    Filter(func(n int) bool { return n > 10 }).
    Take(3).
    Slice()
```

## Creating Pipes

### Of
Create a pipe from a slice.

```go
pipe := pipez.Of([]int{1, 2, 3, 4, 5})
```

## Chain Methods

### Transformation

#### Map
Transform each element.

```go
result := pipez.Of([]int{1, 2, 3}).
    Map(func(n int) int { return n * n }).
    Slice()
// result = []int{1, 4, 9}
```

#### Reverse
Reverse order.

```go
result := pipez.Of([]int{1, 2, 3}).
    Reverse().
    Slice()
// result = []int{3, 2, 1}
```

#### Shuffle
Randomize order.

```go
result := pipez.Of([]int{1, 2, 3, 4, 5}).
    Shuffle().
    Slice()
// result = []int{4, 1, 5, 2, 3} (random)
```

### Filtering

#### Filter
Keep matching elements.

```go
result := pipez.Of([]int{1, 2, 3, 4, 5, 6}).
    Filter(func(n int) bool { return n%2 == 0 }).
    Slice()
// result = []int{2, 4, 6}
```

#### Reject
Remove matching elements.

```go
result := pipez.Of([]int{1, 2, 3, 4, 5, 6}).
    Reject(func(n int) bool { return n%2 == 0 }).
    Slice()
// result = []int{1, 3, 5}
```

#### Take
Take first N.

```go
result := pipez.Of([]int{1, 2, 3, 4, 5}).
    Take(3).
    Slice()
// result = []int{1, 2, 3}
```

#### TakeRight
Take last N.

```go
result := pipez.Of([]int{1, 2, 3, 4, 5}).
    TakeRight(2).
    Slice()
// result = []int{4, 5}
```

#### Drop
Drop first N.

```go
result := pipez.Of([]int{1, 2, 3, 4, 5}).
    Drop(2).
    Slice()
// result = []int{3, 4, 5}
```

#### DropRight
Drop last N.

```go
result := pipez.Of([]int{1, 2, 3, 4, 5}).
    DropRight(2).
    Slice()
// result = []int{1, 2, 3}
```

#### TakeWhile
Take while predicate true.

```go
result := pipez.Of([]int{1, 2, 3, 4, 5}).
    TakeWhile(func(n int) bool { return n < 4 }).
    Slice()
// result = []int{1, 2, 3}
```

#### TakeRightWhile
Take from right while true.

```go
result := pipez.Of([]int{1, 2, 3, 4, 5}).
    TakeRightWhile(func(n int) bool { return n > 2 }).
    Slice()
// result = []int{3, 4, 5}
```

#### DropWhile
Drop while predicate true.

```go
result := pipez.Of([]int{1, 2, 3, 4, 5}).
    DropWhile(func(n int) bool { return n < 4 }).
    Slice()
// result = []int{4, 5}
```

#### DropRightWhile
Drop from right while true.

```go
result := pipez.Of([]int{1, 2, 3, 4, 5}).
    DropRightWhile(func(n int) bool { return n > 3 }).
    Slice()
// result = []int{1, 2, 3}
```

### Combination

#### Concat
Append slices.

```go
result := pipez.Of([]int{1, 2, 3}).
    Concat([]int{4, 5}, []int{6, 7, 8}).
    Slice()
// result = []int{1, 2, 3, 4, 5, 6, 7, 8}
```

#### Zip
Combine with another slice.

```go
result := pipez.Of([]int{1, 2, 3}).
    Zip([]int{10, 20, 30}, func(a, b int) int {
        return a + b
    }).
    Slice()
// result = []int{11, 22, 33}
```

#### Interleave
Interleave with other slices.

```go
result := pipez.Of([]int{1, 4, 7}).
    Interleave([]int{2, 5, 8}, []int{3, 6, 9}).
    Slice()
// result = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
```

### Side Effects

#### Peek
Apply function without changing values.

```go
result := pipez.Of([]int{1, 2, 3}).
    Peek(func(n int) { fmt.Printf("Processing %d\n", n) }).
    Map(func(n int) int { return n * 2 }).
    Slice()
// Prints: Processing 1, Processing 2, Processing 3
// result = []int{2, 4, 6}
```

## Terminal Operations

### Slice
Get the underlying slice.

```go
underlying := pipe.Slice()
```

### Head
Get first element.

```go
first, err := pipe.Head()
// Returns error if empty
```

### Last
Get last element.

```go
last, err := pipe.Last()
// Returns error if empty
```

### Count
Get length.

```go
n := pipe.Count()
```

### Nth
Get element at index.

```go
elem := pipe.Nth(2)  // 3rd element (0-indexed)
```

### Every / Some / None
Check predicates.

```go
allPositive := pipe.Every(func(n int) bool { return n > 0 })
hasEven := pipe.Some(func(n int) bool { return n%2 == 0 })
noNegatives := pipe.None(func(n int) bool { return n < 0 })
```

### Fold / FoldRight
Reduce to value.

```go
sum := pipe.Fold(func(acc, n int) int { return acc + n }, 0)
```

### Partition
Split into two slices.

```go
evens, odds := pipe.Partition(func(n int) bool { return n%2 == 0 })
```

### Compact
Remove consecutive duplicates.

```go
result := pipez.Of([]int{1, 1, 2, 2, 2, 3}).
    Compact(func(a, b int) bool { return a == b }).
    Slice()
// result = []int{1, 2, 3}
```

### Sample
Get random sample.

```go
result := pipez.Of([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
    Sample(3).
    Slice()
// result = 3 random elements
```

### SortFunc
Sort with custom comparator.

```go
result := pipez.Of([]int{3, 1, 4, 1, 5}).
    SortFunc(func(a, b int) bool { return a < b }).
    Slice()
// result = []int{1, 1, 3, 4, 5}
```

## Complex Example

```go
// Process user data
type User struct {
    Name string
    Age  int
}

users := []User{
    {"Alice", 30},
    {"Bob", 25},
    {"Charlie", 35},
    {"David", 20},
}

// Get names of adults, sorted, take top 3
adultNames := pipez.Of(users).
    Filter(func(u User) bool { return u.Age >= 18 }).
    SortFunc(func(a, b User) bool { return a.Age > b.Age }).
    Map(func(u User) string { return u.Name }).
    Take(3).
    Slice()

// adultNames = []string{"Charlie", "Alice", "Bob"}
```

## Performance

`Pipe[A]` is a type alias for `[]A`, so there's **zero overhead**:

```go
type Pipe[A any] []A

// Direct conversion
var slice []int = []int{1, 2, 3}
pipe := pipez.Of(slice)  // Just a cast
back := pipe.Slice()       // Just a cast
```

## When to Use

**Use pipez when:**
- Building complex data pipelines
- Readability is more important than raw performance
- You want to chain 3+ operations

**Use slicez when:**
- You need maximum performance
- Doing single operations
- Working in tight loops

## See Also

- [slicez](../slicez/) - Underlying function implementations
- [chanz](../chanz/) - For asynchronous pipelines
