# Henry

> A collection of nice to have generic function for slices, maps, channels and numbers

[![GoDoc](https://godoc.org/github.com/modfin/henry?status.svg)](https://pkg.go.dev/github.com/modfin/henry)
[![Go report](https://goreportcard.com/badge/github.com/modfin/henry)](https://goreportcard.com/report/github.com/modfin/henry)

## Install

```bash
go get github.com/modfin/henry/...
```

### Usage

Import the part of henry you are interested of using

```go
import (
"github.com/modfin/henry/chanz"
"github.com/modfin/henry/mapz"
"github.com/modfin/henry/numberz"
"github.com/modfin/henry/slicez"
)
```

Then use the functions in the libraries such as

```go 
pets := slicez.Sort([]string{"Dog", "Lizard", "Cat"}) 
// []string{"Cat", ,"Dog", "Lizard"}

```

## Slicez

### Clone

Produces a copy of a given slice

```go 
s := []int{1,2,3}
clone := slicez.Clone[int](s)
// []int{1,2,3}
```

### Compact

Removes consecutive duplicates from a slice

```go 
s := []int{1, 1, 2, 3, 3}
slicez.Compact[int](s)
// []int{1,2,3}
```

### CompactFunc

Removes consecutive duplicates from a slice using a function for determine equality

```go 
s := []rune("Alot    of  white  spaces")
slicez.CompactFunc[rune](s, func(a, b rune) {
     return a == ' ' && a == b
})
// "Alot of white spaces"
```

### Compare

Compares two slices for equality

```go
s1 := []int{1, 2, 3}
s2 := []int{1, 2, 3}
slicez.Compare[int](s1, s2)
// 0
```

### CompareFunc

Compares two slices for equality with supplied func

```go
s1 := []int{1, 2, 3}
s2 := []int{4, 5, 6}
slicez.CompareFunc[int](s1, s2, func (a, b int) int{
return a%4 - b%4
})
// 0
```

### Complement

Returns the complement of two slices

```go 
a := []int{1, 2, 3}
b := []int{3, 2, 5, 5, 6, 1}
slicez.Complement[int](a, b)
// []int{5, 6}
```

### ComplementBy

### Concat

Concatenates slices into a new slice

```go 
a := []int{1, 2, 3}
b := []int{4, 5, 6}
b := []int{7, 8, 9}
slicez.Concat[int](a, b, c)
// []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
```

### Contains

Return true if an element is present in the slice

```go 
slicez.Contains[int]([]int{1, 2, 3, 4, 5}, 3)
// true
```

### ContainsFunc

Returns true if the function returns true.

```go 
slicez.ContainsFunc[int]([]int{4, 5, 6, 7}, func(i int){ 
    return i % 4 == 3
})
// true
```

### Cut

Cuts a slice into two parts

```go 
slicez.Cut[int]([]int{1,2,3,4,5}, 3)
// []int{1,2}, []int{4,5}, true
```

### CutFunc

Cuts a slice into two parts

```go 
slicez.CutFunc[int]([]int{1,2,3,4,5}, func(i int){ return i == 3}) 
// []int{1,2}, []int{4,5}, true
```

### Difference

Returns the difference between slices

```go 
a := []int{1,2,3}
b := []int{2,3,4}
c := []int{2,3,5}
slicez.Difference[int](a, b, c) 
// []int{1,4,5}
```

### DifferenceBy

### Drop

Drops the first N elements

```go
slicez.Drop[int]([]int{1, 2, 3, 4, 5}, 2)
// []int{3,4,5}
```

### DropRight

Drops the last N elements

```go
slicez.DropRight[int]([]int{1, 2, 3, 4, 5}, 2)
// []int{1,2,3}
```

### DropWhile

Drops elements from the left until function returns false

```go
slicez.DropWhile[int]([]int{1, 2, 3, 4, 5}, func (i int) {return i < 3})
// []int{3,4,5}
```

### DropRightWhile

Drops elements from the right until function returns false

```go
slicez.DropRightWhile[int]([]int{1, 2, 3, 4, 5}, func (i int) {return i > 3})
// []int{1,2,3}
```

### Each

Applies a function to each element of a slice

```go
slicez.Each[int]([]int{1, 2, 3}, func (i int) { fmt.Print(i) })
// 123
```

### Equal

Returns true if two slices are identical

```go
s1 := []int{1, 2, 3}
s2 := []int{1, 2, 3}
slicez.Equal[int](s1, s2)
// true
```

### EqualFunc

Returns true if two slices are identical, given an equality function

```go
s1 := []int{1, 2, 3}
s2 := []int{4, 5, 6}
slicez.EqualFunc[int, int](s1, s2, func (a, b int) int{
return a%4 - b%4
})
// true
```

### Every

Returns true if every element matches the given value

```go
s1 := []int{1, 1, 1}
slicez.Every[int](s1, 1)
// true
```

### EveryFunc

### Every

Returns true if every element matches the given value using the equality function

```go
s1 := []int{0, 3, 6}
slicez.EveryFunc[int](s1, func (i int) { return i % 3 == 0})
// true
```

### Filter

Filters a slice to contain things we are looking for

```go
s1 := []int{1, 2, 3, 4}
slicez.Filter[int](s1, func (i int) { return i % 2 == 0})
// []int{2,4}
```

### Find

Find returns the first instance of an object where the function returns true

```go
s1 := []int{1, 2, 3, 4, 5,6, 7}
slicez.Find[int](s1, func (i int) { return i % 3 == 0})
// 3
```

### FindLast

FindLast returns the last instance of an object where the function returns true

```go
s1 := []int{1, 2, 3, 4, 5,6, 7}
slicez.FindLast[int](s1, func (i int) { return i % 3 == 0})
// 6
```

### FlatMap

Takes a slice, expands every element into a slice and flattens it to a single slice

```go 
s := []string{"a b c", "d e f"}

slicez.FlatMap(s, func(e string) []string {
    return strings.Split(e, " ")
})
// []string{"a","b","c","d","e","f"}
```

### Flatten

Flattens a nested slice

```go 
s := [][]string{{"a", "b"}, {"c","d"}}
slicez.Flatten(s)
// []string{"a","b","c","d"}
```

### Fold

Folds a slice into a value from the left (aka reduce)

```go 
s := []string{"a", "b", "c","d"}
slicez.Fold[string, string](s, func(acc string, str string) string { return acc + str}, ">")
// ">abcd"
```

### FoldRight

Folds a slice into a value from the right (aka reduce)

```go 
s := []string{"a", "b", "c","d"}
slicez.FoldRight[string, string](s, func(acc string, str string) string { return acc + str}, ">")
// ">dcba"
```

### GroupBy

Groups elements in a slice into a map

```go 
s = []int{0,1,2,3}
slicez.GroupBy(s, func(e int) int {return i % 2})
// map[int][]int{0: [0,2], 1:[1,2]}
```

### Head

Returns the first element of a slice if present

```go 
slicez.Head([]int{1,2,3})
// 1, nil
```

### Index

Returns the index of the first occurrence of an element

```go  
slicez.Index([]int{1,2,3}, 2)
// 1
```

### IndexFunc

Returns the index of the first occurrence of an element using a function

```go  
slicez.IndexFunc([]int{1,2,3}, func(i int) bool {i % 2 == 1})
// 0
```

### Intersection

Returns the intersection of slices

```go 
a := []int{1,2,3,4}
b := []int{3,4,5,6}
c := []int{3,4,8,9}
slicez.Intersection(a,b,c)
// []int{3,4}
```

### IntersectionBy

### Intersection

Returns the intersection of slices

```go 
a := []int{0,1}
b := []int{4,2}
c := []int{8,3}
slicez.IntersectionBy(func(i int) int {
    return i % 4
} a,b,c)
// []int{0,4,8}
```

### Join

Joining a 2d slice into 1d slice using a glue

```go 
s = [][]string{{"hello", " ", "world"}, {"or", " ", "something"}}
slicez.Join(s, []string{" "})
// []string{"hello", " ", "world", " ", "or", s" ", "something"}
```

### KeyBy

Returns a map with the slice elements in it, using the by function to determine key

```go 
s = []int{1,2,3,4}
slicez.KeyBy(s, func(i int) int { return i % 3 })
// map[int]int{0: 3, 1: 1, 2: 2}
```

### Last

Returns the last element in a slice, or an error if len(s) == 0

```go 
slicez.Last([]int{1,2,3})
// 3, nil
```

### LastIndex

Finds the last index of a needle, or -1 if not present

```go 
slicez.LastIndex([]int{1,1,2,1,3}, 1)
// 3
```

### LastIndexFunc

### LastIndex

Finds the last index of a func needle, or -1 if not present

```go 
slicez.LastIndex([]int{1,2,3,4,5}, func(i int) bool { return i % 3 == 1})
// 2
```

### Map

Map values in a slice producing a new one

```go 
s := []int{1,2,3}
slicez.Map(s, func(i int) string { return fmt.Sprint(i)})
// []string{"1","2","3"}
```

### Max

Returns the maximum value of a slice

```go 
s := []int{1,2,5,4}
slicez.Max(s...)
// 5
```

### Min

Returns the minimum value of a slice

```go 
s := []int{1,2,5,0, 4}
slicez.Min(s...)
// 0
```

### None
Returns true if no element match the needle 

```go 
s := []int{1,2,3}
slicez.None(s, 0)
// true
```


### NoneFunc
Returns true if no element returns true from the function

```go 
s := []int{1,2,3}
slicez.NoneFunc(s, func(i int) bool { return i < 1 })
// true
```

### Nth
Returns the N:th element in a slice, zero value if empty and regards the slice as a modulo group

```go 
s := []int{1,2,3}
slicez.Nth(s, 1)
// 2
slicez.Nth(s, 3)
// 1
slicez.Nth(s, -1)
// 3
```


### Partition
Returns two slices which represents the partitions

```go 
s := []int{1,2,3,4}
slicez.Partition(s, func(i int) bool { return i % 2 == 0})
// []int{2,4}, []int{1,3}
```


### Reject
Reject is the complement to Filter and excludes items
```go 
s := []int{1,2,3,4}
slicez.Reject(s, func(i int) bool { return i % 2 == 0})
// []int{1,3}
```


### Reverse
Reverses a slice

```go 
s := []int{1,2,3}
slicez.Reverse(s)
//[]int{3,2,1}
```

### Sample
Returns a random sample of size N from the slice

```go 
s := []int{1,2,3,4,5,6,7,8}
slicez.Sample(s, 2)
//[]int{8,3}
```

### Search

### Shuffle
Returns a shuffled version of the slice

```go 
s := []int{1,2,3,4,5}
slicez.Shuffle(s)
//[]int{3,1,4,5,3}
```

### Some
Returns true there exist an element in the slice that is equal to the needle, an alias for Contains

```go 
s := []int{1,2,3,4,5}
slicez.Some(s, 4)
//true
```


### SomeFunc
Returns true if there is an element in the slice for which the predicate function returns true

```go 
s := []int{1,2,3,4,5}
slicez.SomeFunc(s, func(i int) bool { return i > 4})
//true
```


### Sort
Sorts a slice
```go 
s := []int{3,2,1}
slicez.Sort(s)
//[]int{1,2,3}
```

### SortFunc
Sorts a slice with a comparator
```go 
s := []int{1,2,3}
slicez.SortFunc(s, func(a, b int) bool { return b < a })
//[]int{3,2,1}
```

### Tail
Returns the tail of a slice

```go 
s := []int{1,2,3}
slicez.Tail(s)
// []int{2,3}
```

### Take
Returns the N first element of a slice
```go 
s := []int{1,2,3,4}
slicez.Take(s, 2)
// []int{1,2}
```

### TakeRight
Returns the N last element of a slice
```go 
s := []int{1,2,3,4}
slicez.TakeRight(s, 2)
// []int{3, 4}
```

### TakeWhile
Returns the first element of a slice that as long as function returns t
```go 
s := []int{1,2,3,4}
slicez.TakeRight(s, func(i int) bool { return i < 3})
// []int{1, 2}
```

### TakeRightWhile
Returns the last element of a slice that as long as function returns t
```go 
s := []int{1,2,3,4}
slicez.TakeWhileRight(s, func(i int) bool { return i > 2})
// []int{3, 4}
```


### Union
Returs the union of a slices

```go 
a := []int{1,2,3}
b := []int{3,4,5}
slicez.Union(a, b)
// []int{1,2,3,4,5}
```

### UnionBy
Returs the union of a slices using a function for equality
```go 
a := []int{1,5}
b := []int{2,4}
slicez.UnionBy(func(i int) bool { return i % 2 == 0 } a, b)
// []int{1,2}
```

### Uniq
Returns a slice of uniq elements
```go 
a := []int{1,2,3,1,3,4}
slicez.Uniq(a)
// []int{1,2,3,4}
```

### UniqBy
Returns a slice of uniq elements, where equality is determined through the function
```go 
a := []int{1,2,3,1,3,4}
slicez.UniqBy(a, func(i int) bool { return i % 2 == 0 })
// []int{1,2}
```

### Unzip
Takes a slice and unzips it into two slices

```go 
s := []int{-1,2}
slicez.Unzip(s, func(i int) (bool, int){
    return i > 0, int(Math.Abs(i))
})
// []bool{false, true}, []int{1,2}
```


### Unzip2
Takes a slice and unzips it into three slices
```go 
s := []int{-2,-1,2}
slicez.Unzip(s, func(i int) (bool, bool, int){
    return i > 0, i % 2 == 0, int(Math.Abs(i))
})
// []bool{false, false, true}, []bool{true, false, true}, []int{2, 1,2}
```


### Zip
Takes 2 slices and zips them into one slice

```go 
a := []int{1,2,3}
b := []string{"a","b","c"}
slicez.Zip(a,b, func(i int, s string) string {
    return fmt.Sprint(i, s)
})
// []string{"1a", "2b", "3c"}
```

### Zip2
Takes 3 slices and zips them into one slice

```go 
a := []int{1,2,3}
b := []string{"a","b","c"}
b := []bool{true, false, true}
slicez.Zip(a, b, c, func(i int, s string, b bool) string {
    return fmt.Sprint(b, i, s)
})
// []string{"true1a", "false2b", "true3c"}
```


## Chanz

### Collect

### CollectUntil

### Compact

### Compact1

### CompactN

### CompactUntil

### Concat

### Concat1

### ConcatN

### ConcatUntil

### Drop

### Drop1

### DropAll

### DropN

### DropUntil

### DropWhile

### DropWhile1

### DropWhileN

### DropWhileUntil

### FanOut

### FanOut1

### FanOutN

### FanOutUntil

### Filter

### Filter1

### FilterN

### FilterUntil

### Flatten

### Flatten1

### FlattenN

### FlattenUntil

### Generate

### Generate1

### GenerateN

### GenerateUntil

### Map

### Map1

### MapN

### MapUntil

### Merge

### Merge1

### MergeN

### MergeUntil

### Partition

### Partition1

### PartitionN

### PartitionUntil

### Peek

### Peek1

### PeekN

### PeekUntil

### Readers

### SomeDone

### Take

### Take1

### TakeN

### TakeUntil

### TakeWhile

### TakeWhile1

### TakeWhileN

### TakeWhileUntil

### Unzip

### Unzip1

### UnzipN

### UnzipUntil

### Writers

### Zip

### Zip1

### ZipN

### ZipUntil

## Mapz

### Clear

### Clone

### Copy

### DeleteFunc

### DeleteValue

### Equal

### EqualFunc

### Keys

### Merge

### Remap

### Values

## Numberz

### BitAND

### BitOR

### BitXOR

### Corr

### Cov

### FTest

### GCD

### LCM

### LinReg

### MAD

### MapAbs

### MapByte

### MapFloat32

### MapFloat64

### MapInt

### MapInt16

### MapInt32

### MapInt64

### MapInt8

### MapNegate

### MapUInt

### MapUInt16

### MapUInt32

### MapUInt64

### MapUInt8

### Max

### Mean

### Median

### Min

### Mode

### Modes

### Percentile

### R2

### Range

### SNR

### Skew

### StdDev

### StdErr

### Sum

### VAdd

### VDot

### VMul

### VPow

### VSub

### Var

### ZScore

