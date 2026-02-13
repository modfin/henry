# chanz

> Generic channel pipelines and utilities for Go

The `chanz` package provides 50+ functions for building channel pipelines, fan-in/fan-out patterns, and concurrent data processing. All functions support functional options for configuration.

## Quick Reference

**By Category:**
- [Generation](#generation) - Generate, Generator
- [Transformation](#transformation) - Map, Flatten, Zip, Unzip
- [Filtering](#filtering) - Filter, Compact, Take, Drop, Partition
- [Aggregation](#aggregation) - FanIn, Concat, Collect
- [Fan-Out](#fan-out) - FanOut
- [Control Flow](#control-flow) - Done signals, SomeDone, EveryDone
- [Buffering](#buffering) - Buffer, TakeBuffer, DropBuffer, DropAll
- [Channel Types](#channel-types) - Readers, Writers

## Installation

```bash
go get github.com/modfin/henry/chanz
```

## Usage

```go
import "github.com/modfin/henry/chanz"

// Create a pipeline
input := chanz.Generate(1, 2, 3, 4, 5)
doubled := chanz.Map(input, func(n int) int { return n * 2 })
result := chanz.Collect(doubled)
// result = []int{2, 4, 6, 8, 10}
```

## Configuration Options

Most functions accept options:

```go
// Buffer size
ch := chanz.GenerateWith[int](chanz.OpBuffer(100))(1, 2, 3, 4, 5)

// Context cancellation
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
ch := chanz.Map(input, transform, chanz.OpContext(ctx))

// Done channel
done := make(chan struct{})
ch := chanz.Map(input, transform, chanz.OpDone(done))
```

## Function Categories

### Generation

Create channels from data or generators.

#### Generate
Create channel from elements.

```go
ch := chanz.Generate(1, 2, 3, 4, 5)
result := chanz.Collect(ch)
// result = []int{1, 2, 3, 4, 5}
```

#### GenerateWith
Configured generator.

```go
// With 10-element buffer
gen := chanz.GenerateWith[int](chanz.OpBuffer(10))
ch := gen(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
```

#### Generator
Create channel from generator function.

```go
// Generate Fibonacci numbers
fib := chanz.Generator(func(yield func(int)) {
    a, b := 0, 1
    for i := 0; i < 10; i++ {
        yield(a)
        a, b = b, a+b
    }
})
result := chanz.Collect(fib)
// result = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
```

#### GeneratorWith
Configured generator function.

```go
fibGen := chanz.GeneratorWith[int](chanz.OpBuffer(5))
fib := fibGen(func(yield func(int)) {
    yield(0)
    yield(1)
    yield(1)
})
```

### Transformation

Transform channel data.

#### Map
Transform each element.

```go
input := chanz.Generate(1, 2, 3, 4, 5)
doubled := chanz.Map(input, func(n int) int {
    return n * 2
})
result := chanz.Collect(doubled)
// result = []int{2, 4, 6, 8, 10}
```

#### MapWith
Configured mapper.

```go
doubler := chanz.MapWith[int, int](chanz.OpBuffer(10))
input := chanz.Generate(1, 2, 3, 4, 5)
result := chanz.Collect(doubler(input, func(n int) int {
    return n * 2
}))
```

#### Peek
Side-effect without transformation.

```go
input := chanz.Generate(1, 2, 3)
logged := chanz.Peek(input, func(n int) {
    fmt.Printf("Processing: %d\n", n)
})
result := chanz.Collect(logged)
// Prints: Processing: 1, Processing: 2, Processing: 3
// result = []int{1, 2, 3}
```

#### Flatten
Flatten channel of slices.

```go
input := chanz.Generate([]int{1, 2, 3}, []int{4, 5, 6})
flat := chanz.Flatten(input)
result := chanz.Collect(flat)
// result = []int{1, 2, 3, 4, 5, 6}
```

#### Zip
Combine two channels.

```go
nums := chanz.Generate(1, 2, 3)
strs := chanz.Generate("a", "b", "c")
zipped := chanz.Zip(nums, strs, func(n int, s string) string {
    return fmt.Sprintf("%d:%s", n, s)
})
result := chanz.Collect(zipped)
// result = []string{"1:a", "2:b", "3:c"}
```

#### Unzip
Split channel into two.

```go	type Pair struct{ X, Y int }
input := chanz.Generate(Pair{1, 10}, Pair{2, 20}, Pair{3, 30})
xs, ys := chanz.Unzip(input, func(p Pair) (int, int) {
    return p.X, p.Y
})
// xs receives 1, 2, 3
// ys receives 10, 20, 30
```

### Filtering

Select or skip elements.

#### Filter
Keep matching elements.

```go
input := chanz.Generate(1, 2, 3, 4, 5, 6)
evens := chanz.Filter(input, func(n int) bool {
    return n%2 == 0
})
result := chanz.Collect(evens)
// result = []int{2, 4, 6}
```

#### Compact
Remove consecutive duplicates.

```go
input := chanz.Generate(1, 1, 2, 2, 2, 3, 3)
compacted := chanz.Compact(input, func(a, b int) bool {
    return a == b
})
result := chanz.Collect(compacted)
// result = []int{1, 2, 3}
```

#### Take
Take first N elements.

```go
input := chanz.Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
first3 := chanz.Take(input, 3)
result := chanz.Collect(first3)
// result = []int{1, 2, 3}
```

#### TakeWhile
Take while predicate true.

```go
input := chanz.Generate(1, 2, 3, 4, 5, 1, 2)
ascending := chanz.TakeWhile(input, func(n int) bool {
    return n < 4
})
result := chanz.Collect(ascending)
// result = []int{1, 2, 3}
```

#### Drop
Drop first N elements.

```go
input := chanz.Generate(1, 2, 3, 4, 5, 6)
rest := chanz.Drop(input, 3)
result := chanz.Collect(rest)
// result = []int{4, 5, 6}
```

#### DropWhile
Drop while predicate true.

```go
input := chanz.Generate(1, 2, 3, 4, 5)
from4 := chanz.DropWhile(input, func(n int) bool {
    return n < 4
})
result := chanz.Collect(from4)
// result = []int{4, 5}
```

#### Partition
Split into two channels.

```go
input := chanz.Generate(1, 2, 3, 4, 5, 6)
evens, odds := chanz.Partition(input, func(n int) bool {
    return n%2 == 0
})
// evens receives 2, 4, 6
// odds receives 1, 3, 5
```

### Aggregation

Combine multiple channels.

#### FanIn
Merge channels concurrently.

```go
ch1 := chanz.Generate(1, 2, 3)
ch2 := chanz.Generate(4, 5, 6)
ch3 := chanz.Generate(7, 8, 9)

merged := chanz.FanIn(ch1, ch2, ch3)
result := chanz.Collect(merged)
// result contains all 9 numbers (order non-deterministic)
```

#### Concat
Concatenate channels sequentially.

```go
ch1 := chanz.Generate(1, 2, 3)
ch2 := chanz.Generate(4, 5, 6)
ch3 := chanz.Generate(7, 8, 9)

combined := chanz.Concat(ch1, ch2, ch3)
result := chanz.Collect(combined)
// result = []int{1, 2, 3, 4, 5, 6, 7, 8, 9} (order preserved)
```

#### Collect
Read all elements into slice.

```go
input := chanz.Generate(1, 2, 3, 4, 5)
result := chanz.Collect(input)
// result = []int{1, 2, 3, 4, 5}
```

### Fan-Out

Distribute to multiple channels.

#### FanOut
Broadcast to multiple channels.

```go
input := chanz.Generate(1, 2, 3, 4, 5)
outputs := chanz.FanOut(input, 3, chanz.OpBuffer(1))

// outputs[0], outputs[1], outputs[2] all receive: 1, 2, 3, 4, 5

// Process in parallel
for i, ch := range outputs {
    go func(id int, c <-chan int) {
        for n := range c {
            fmt.Printf("Worker %d got %d\n", id, n)
        }
    }(i, ch)
}
```

### Control Flow

Signal coordination and cancellation.

#### Done
Convert any channel to done signal.

```go
work := make(chan int)
done := chanz.Done(work)

// Close work channel to signal done
close(work)
<-done // Unblocks when work is closed
```

#### SomeDone
Close when ANY input closes.

```go
done1 := make(chan struct{})
done2 := make(chan struct{})

done := chanz.SomeDone(done1, done2)

// Close either channel
close(done1)
<-done // Unblocks immediately
```

#### EveryDone
Close when ALL inputs close.

```go
done1 := make(chan struct{})
done2 := make(chan struct{})

done := chanz.EveryDone(done1, done2)

// Must close both
close(done1)
close(done2)
<-done // Unblocks now
```

### Buffering

Buffer management utilities.

#### Buffer
Collect N elements with done support.

```go
input := chanz.Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
batch, more := chanz.Buffer(3, input)
// batch = []int{1, 2, 3}, more = true (channel still open)

batch, more = chanz.Buffer(100, input)
// batch = []int{4, 5, 6, 7, 8, 9, 10}, more = false (channel closed)
```

#### TakeBuffer
Non-blocking buffer read.

```go
ch := make(chan int, 10)
ch <- 1
ch <- 2
ch <- 3

buffered := chanz.TakeBuffer(ch)
// buffered = []int{1, 2, 3}
// ch still has space for 7 more
```

#### DropBuffer
Non-blocking buffer clear.

```go
ch := make(chan int, 10)
ch <- 1
ch <- 2
ch <- 3

chanz.DropBuffer(ch, false)
// ch is now empty (0 buffered items)
```

#### DropAll
Consume until closed.

```go
input := chanz.Generate(1, 2, 3, 4, 5)

// Synchronous - blocks until channel closed
chanz.DropAll(input, false)

// Asynchronous - returns immediately, drains in background
chanz.DropAll(input, true)
```

### Channel Types

Type conversions for safety.

#### Readers
Convert to read-only channels.

```go
chans := []chan int{make(chan int), make(chan int)}
readers := chanz.Readers(chans...)
// readers is []<-chan int
```

#### Writers
Convert to write-only channels.

```go
chans := []chan int{make(chan int), make(chan int)}
writers := chanz.Writers(chans...)
// writers is []chan<- int
```

### Write/Read Modes

Flexible I/O operations.

#### WriteTo
Create writer with mode.

```go
ch := make(chan int, 1)

// Synchronous (blocks)
writeSync := chanz.WriteTo[int](ch, chanz.WriteSync)
writeSync(42) // Blocks until written

// Asynchronous (goroutine)
writeAsync := chanz.WriteTo[int](ch, chanz.WriteAync)
writeAsync(42) // Returns immediately

// Non-blocking (only if space)
writeIfFree := chanz.WriteTo[int](ch, chanz.WriteIfFree)
writeIfFree(42) // Only writes if buffer has space
```

#### ReadFrom
Create reader with mode.

```go
ch := make(chan int)
ch <- 42

// Synchronous (blocks)
readSync := chanz.ReadFrom[int](ch, chanz.ReadWait)
val, ok := readSync() // Blocks, returns (42, true)

// Non-blocking
readNow := chanz.ReadFrom[int](ch, chanz.ReadIfWaiting)
val, ok := readNow() // Returns immediately
```

## Common Patterns

### Pipeline Pattern

```go
// Build a processing pipeline
input := chanz.Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

processed := chanz.Map(input, func(n int) int {
    return n * n
})

filtered := chanz.Filter(processed, func(n int) bool {
    return n > 10
})

result := chanz.Collect(filtered)
// result = []int{16, 25, 36, 49, 64, 81, 100}
```

### Worker Pool

```go
// Generate work
work := chanz.Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

// Fan out to 3 workers
workers := chanz.FanOut(work, 3)

// Process and collect results
var wg sync.WaitGroup
results := make(chan int, 10)

for i, worker := range workers {
    wg.Add(1)
    go func(id int, ch <-chan int) {
        defer wg.Done()
        for n := range ch {
            // Simulate work
            time.Sleep(100 * time.Millisecond)
            results <- n * 2
            fmt.Printf("Worker %d processed %d\n", id, n)
        }
    }(i, worker)
}

// Close results when done
go func() {
    wg.Wait()
    close(results)
}()

// Collect all results
final := chanz.Collect(results)
```

### Timeout Handling

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

input := chanz.Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

// Slow transformation
slow := chanz.Map(input, func(n int) int {
    time.Sleep(1 * time.Second)
    return n * 2
}, chanz.OpContext(ctx))

result := chanz.Collect(slow)
// Will stop early if timeout exceeded
```

## Performance Notes

- **Goroutine per channel**: Map, Filter, etc. spawn goroutines
- **Buffered channels**: Use `OpBuffer(n)` to improve throughput
- **Backpressure**: FanOut waits for all outputs to consume
- **Clean shutdown**: Always close input channels to signal completion
- **Context cancellation**: Use `OpContext()` for graceful shutdown

## See Also

- [pipez](../pipez/) - Fluent API for synchronous operations
- [slicez](../slicez/) - Batch operations on collected data
