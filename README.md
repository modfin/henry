# Henry

> Generic utility functions for Go - functional programming made easy

[![GoDoc](https://godoc.org/github.com/modfin/henry?status.svg)](https://pkg.go.dev/github.com/modfin/henry)
[![Go Report Card](https://goreportcard.com/badge/github.com/modfin/henry?cache)](https://goreportcard.com/report/github.com/modfin/henry)

Henry is a collection of generic utility functions for Go 1.18+, providing functional programming primitives for slices, maps, and channels. It offers both standalone functions and fluent APIs to make your code more expressive and maintainable.

## Why Henry?

- **Zero-allocation overhead** - Type aliases where possible, no wrapper structs
- **Generic-first design** - Leverages Go 1.18+ generics for type safety
- **Consistent API** - Similar patterns across all packages
- **Well-documented** - Every function has clear examples
- **Production-ready** - Comprehensive test coverage (95%+)

## Quick Start

```bash
go get github.com/modfin/henry/...
```

```go
import (
    "github.com/modfin/henry/slicez"
    "github.com/modfin/henry/mapz"
    "github.com/modfin/henry/chanz"
)

// Transform data
numbers := []int{1, 2, 3, 4, 5}
doubled := slicez.Map(numbers, func(n int) int { return n * 2 })
// doubled = []int{2, 4, 6, 8, 10}

// Filter collections
evens := slicez.Filter(numbers, func(n int) bool { return n%2 == 0 })
// evens = []int{2, 4}

// Work with maps
scores := map[string]int{"Alice": 85, "Bob": 92}
names := mapz.Keys(scores)
// names = []string{"Alice", "Bob"} (order not guaranteed)

// Process streams
input := chanz.Generate(1, 2, 3, 4, 5)
doubledCh := chanz.Map(input, func(n int) int { return n * 2 })
result := chanz.Collect(doubledCh)
// result = []int{2, 4, 6, 8, 10}
```

## Package Overview

| Package | Description | Key Features |
|---------|-------------|--------------|
| **[slicez](./slicez/)** | Slice operations | Map, Filter, Reduce, Sort, Set operations |
| **[mapz](./mapz/)** | Map operations | Keys/Values, Filter, Merge, Transform |
| **[chanz](./chanz/)** | Channel pipelines | FanIn/FanOut, Map, Filter, Done signals |
| **[setz](./setz/)** | Set data structure | Union, Intersection, Difference |
| **[pipez](./pipez/)** | Fluent API | Method chaining for slice operations |
| **[compare](./compare/)** | Comparison utilities | Less, Greater, Between, Clamp |
| **[mon](./mon/)** | Monadic types | Result, Option types for error handling |

## Common Patterns

### Functional Pipeline

```go
// Traditional imperative approach
var evenSquares []int
for _, n := range numbers {
    if n%2 == 0 {
        evenSquares = append(evenSquares, n*n)
    }
}

// Functional approach with Henry
evenSquares := pipez.Of(numbers).
    Filter(func(n int) bool { return n%2 == 0 }).
    Map(func(n int) int { return n * n }).
    Slice()
```

### Safe Error Handling

```go
// Using the Result monad from mon package
import "github.com/modfin/henry/mon"

urls := []string{"https://example.com", "https://github.com", "bad url"}

// Parse URLs safely
parsed := slicez.Map(urls, func(s string) mon.Result[*url.URL] {
    u, err := url.Parse(s)
    return mon.TupleToResult(u, err)
})

// Extract valid URLs and errors
validURLs, errs := mon.Partition(parsed)
```

### Concurrent Processing

```go
// Generate work items
work := chanz.Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

// Process in parallel (fan-out)
workers := chanz.FanOut(work, 3)

// Process each worker channel
for i, worker := range workers {
    go func(id int, ch <-chan int) {
        for n := range ch {
            result := process(n)
            fmt.Printf("Worker %d processed %d\n", id, result)
        }
    }(i, worker)
}
```

## Performance

Henry is designed with performance in mind:

- **Pre-allocated slices** - Functions like `Map`, `Filter`, and `Union` pre-calculate capacity
- **Zero-overhead types** - `Pipe` and `Set` are type aliases, not wrapper structs
- **Efficient algorithms** - Uses Fisher-Yates shuffle, swap-to-end sampling
- **In-place operations** - Some functions (clearly marked) mutate for performance

```go
// Pre-allocation example
data := make([]int, 10000)
// Map pre-allocates result slice with same capacity
result := slicez.Map(data, transform) // No reallocations!
```

## Comparison with Standard Library

| Task | Standard Library | Henry |
|------|-----------------|-------|
| Sort slice | `sort.Slice()` | `slicez.Sort()` or `slicez.SortBy()` |
| Filter slice | Manual loop | `slicez.Filter()` |
| Map slice | Manual loop | `slicez.Map()` |
| Check contains | Manual loop | `slicez.Contains()` |
| Merge maps | Manual loop | `mapz.Merge()` |
| Channel transform | Manual goroutine | `chanz.Map()` |

## Related Projects

- [golang.org/x/exp](https://github.com/golang/exp) - Official experimental packages
- [samber/lo](https://github.com/samber/lo) - Lodash-style library
- [thoas/go-funk](https://github.com/thoas/go-funk) - Functional utilities

## Contributing

Contributions welcome! Please ensure:
- Tests pass: `go test ./...`
- Code is formatted: `go fmt ./...`
- Documentation is updated for new functions

## License

MIT License - see LICENSE file for details.
