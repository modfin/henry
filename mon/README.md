# mon

> Monadic types for Go - Result and Option

The `mon` package provides functional programming monads: `Result[T]` for error handling and `Option[T]` for optional values. These types make error handling explicit and composable.

## Quick Reference

**By Category:**
- [Result Type](#result-type) - Error handling monad
- [Option Type](#option-type) - Optional values

## Installation

```bash
go get github.com/modfin/henry/mon
```

## Result Type

`Result[T]` represents either a success value or an error. It forces explicit error handling.

### Creating Results

#### Ok
Create a success result.

```go
r := mon.Ok(42)
```

#### Err
Create an error result.

```go
r := mon.Err[int](errors.New("something went wrong"))
```

#### TupleToResult / From
Convert (value, error) tuple to Result.

```go
// From function that returns (T, error)
val, err := strconv.Atoi("42")
r := mon.TupleToResult(val, err)
// r is Ok(42) if err is nil, otherwise Err(err)

// Shorthand
r := mon.From(strconv.Atoi("42"))
```

### Working with Results

#### IsOk / IsErr
Check state.

```go
r := mon.Ok(42)
r.IsOk()   // true
r.IsErr()  // false
```

#### Unwrap
Get value or panic on error.

```go
val := r.Unwrap()  // 42

// Error case - panics!
bad := mon.Err[int](errors.New("oops"))
bad.Unwrap()  // panic!
```

#### UnwrapOr
Get value or default.

```go
val := r.UnwrapOr(0)        // 42
bad.UnwrapOr(0)             // 0 (default)
```

#### UnwrapOrElse
Get value or compute default.

```go
val := bad.UnwrapOrElse(func() int {
    return expensiveFallback()
})
```

#### Expect
Unwrap with custom panic message.

```go
val := r.Expect("should have a value")  // 42
bad.Expect("should have a value")       // panics with message
```

### Transforming Results

#### Map
Transform success value.

```go
r := mon.Ok(5)
doubled := r.Map(func(n int) int {
    return n * 2
})
// doubled = Ok(10)

// Error results pass through unchanged
bad := mon.Err[int](errors.New("oops"))
result := bad.Map(func(n int) int { return n * 2 })
// result still contains the error
```

#### MapErr
Transform error.

```go
bad := mon.Err[int](errors.New("original"))
wrapped := bad.MapErr(func(err error) error {
    return fmt.Errorf("wrapped: %w", err)
})
```

#### FlatMap / Bind
Chain operations that return Results.

```go
r := mon.Ok(5)
result := r.FlatMap(func(n int) mon.Result[int] {
    if n < 0 {
        return mon.Err[int](errors.New("negative"))
    }
    return mon.Ok(n * 2)
})
// result = Ok(10)
```

#### Or
Use alternative if error.

```go
r1 := mon.Err[int](errors.New("failed"))
r2 := mon.Ok(42)
result := r1.Or(r2)
// result = Ok(42)
```

#### OrElse
Compute alternative if error.

```go
result := r1.OrElse(func() mon.Result[int] {
    return fallbackOperation()
})
```

### Collections of Results

#### Partition
Split into successes and failures.

```go
results := []mon.Result[int]{
    mon.Ok(1),
    mon.Err[int](errors.New("bad")),
    mon.Ok(3),
    mon.Err[int](errors.New("worse")),
}

oks, errs := mon.Partition(results)
// oks = []int{1, 3}
// errs = []error{error1, error2}
```

#### Unwrap (collection)
Extract values or return first error.

```go
results := []mon.Result[int]{
    mon.Ok(1),
    mon.Ok(2),
    mon.Ok(3),
}
vals, err := mon.Unwrap(results)
// vals = []int{1, 2, 3}, err = nil

// With an error
badResults := []mon.Result[int]{
    mon.Ok(1),
    mon.Err[int](errors.New("failed")),
    mon.Ok(3),
}
vals, err := mon.Unwrap(badResults)
// vals = nil, err = error
```

## Option Type

`Option[T]` represents either a value or nothing (like nullable types).

### Creating Options

#### Some
Create with a value.

```go
o := mon.Some(42)
```

#### None
Create empty.

```go
o := mon.None[int]()
```

#### FromPtr
Create from pointer.

```go
val := 42
ptr := &val
o := mon.FromPtr(ptr)   // Some(42)

var nilPtr *int
o = mon.FromPtr(nilPtr) // None[int]()
```

### Working with Options

#### IsSome / IsNone
Check state.

```go
o := mon.Some(42)
o.IsSome()  // true
o.IsNone()  // false
```

#### Unwrap
Get value or panic.

```go
val := o.Unwrap()  // 42

// None case
none := mon.None[int]()
none.Unwrap()  // panic!
```

#### UnwrapOr / UnwrapOrElse
Get value or default.

```go
o.UnwrapOr(0)           // 42
mon.None[int]().UnwrapOr(0)  // 0

// With computation
mon.None[int]().UnwrapOrElse(func() int {
    return expensiveDefault()
})
```

#### Expect
Unwrap with message.

```go
val := o.Expect("should have value")
```

### Transforming Options

#### Map
Transform value if present.

```go
o := mon.Some(5)
doubled := o.Map(func(n int) int {
    return n * 2
})
// doubled = Some(10)

// None passes through
mon.None[int]().Map(func(n int) int { return n * 2 })
// still None
```

#### FlatMap / Bind
Chain optional operations.

```go
o := mon.Some(5)
result := o.FlatMap(func(n int) mon.Option[int] {
    if n < 0 {
        return mon.None[int]()
    }
    return mon.Some(n * 2)
})
// result = Some(10)
```

#### Filter
Convert to None if predicate fails.

```go
o := mon.Some(5)
even := o.Filter(func(n int) bool {
    return n%2 == 0
})
// even = None[int]() (5 is odd)
```

#### Or / OrElse
Provide alternative.

```go
none := mon.None[int]()
alt := mon.Some(42)
result := none.Or(alt)
// result = Some(42)
```

## Common Patterns

### Error Handling in Pipelines

```go
import (
    "github.com/modfin/henry/mon"
    "github.com/modfin/henry/slicez"
)

// Parse URLs safely
func parseURLs(urls []string) ([]*url.URL, error) {
    results := slicez.Map(urls, func(s string) mon.Result[*url.URL] {
        u, err := url.Parse(s)
        return mon.From(u, err)
    })
    
    return mon.Unwrap(results)
}

// Usage
urls := []string{
    "https://example.com",
    "https://github.com",
    "bad url",
}

parsed, err := parseURLs(urls)
// parsed contains valid URLs, err is first parse error
```

### Optional Values

```go
func findUser(id int) mon.Option[User] {
    user, err := db.GetUser(id)
    if err != nil {
        return mon.None[User]()
    }
    return mon.Some(user)
}

// Usage
userOpt := findUser(42)
if user, ok := userOpt.UnwrapOr(User{}); ok {
    fmt.Println(user.Name)
}

// Or with default
user := findUser(42).UnwrapOr(guestUser)
```

### Chaining Operations

```go
result := fetchUser(id).
    FlatMap(func(u User) mon.Result[Profile] {
        return fetchProfile(u.ProfileID)
    }).
    Map(func(p Profile) Profile {
        p.LastSeen = time.Now()
        return p
    })

if result.IsErr() {
    log.Printf("Failed: %v", result.Err())
    return
}
profile := result.Unwrap()
```

## Comparison with Standard Error Handling

**Standard Go:**
```go
val1, err := step1()
if err != nil {
    return nil, err
}
val2, err := step2(val1)
if err != nil {
    return nil, err
}
return step3(val2)
```

**With Result:**
```go
return step1().
    FlatMap(step2).
    FlatMap(step3)
```

## When to Use

**Use Result when:**
- Error handling is the primary concern
- You want to defer error checking
- Building complex pipelines

**Use Option when:**
- Nullable values
- Optional configuration
- Caching/memoization

**Use standard Go when:**
- Simple sequential code
- Early returns are needed
- Team prefers idiomatic Go

## See Also

- [slicez](../slicez/) - Use with Map to create collections of Results
- [chanz](../chanz/) - Async pipelines with error handling
