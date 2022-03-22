# Henry

> A collection of nice to have generic function and algorithms for slices, maps and channels

[![GoDoc](https://godoc.org/github.com/modfin/henry?status.svg)](https://pkg.go.dev/github.com/modfin/henry)
[![Go report](https://goreportcard.com/badge/github.com/modfin/henry)](https://goreportcard.com/report/github.com/modfin/henry)

This project came about in the experimentation with go1.18 and generics with helper functions for slices. It now also includes
algorithms and constructs for dealing with channels and maps as well

It is expected to that there might be a lot of these types of libraries floating around after the release of go1.18.
The go team did not include any of these fairly common constructs in this release but instead put some of them into 
the exp package. There for the might be quite a bit of overlap in the coming releases with this and other packages.

Some other work with similar concepts
* https://github.com/golang/exp/tree/master/slices
* https://github.com/golang/exp/tree/master/maps
* https://github.com/samber/lo

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
"github.com/modfin/henry/slicez"
)
```

Then use the functions in the libraries such as

```go 
upperPets := slicez.Map([]string{"Dog", "Lizard", "Cat"}, strings.ToUpper)) 
// []string{"DOG", "LIZARD", "CAT"}

```


#### Error handling
In go errors is made visible and is a core construct for sound code, so we can't simply ignore them.
However, it is not given how to deal with them in a generic functional that makes sense in go and with functions such as Map.
The suggested way of dealing with them is to wrap the result in a result type. This does have some implication in that 
early returns might not be possible and might introduce some extra looping the check the result.

**Example**
```go 
package main

import (
	"fmt"
	"github.com/modfin/henry/exp/result"
	"github.com/modfin/henry/slicez"
	"net/url"
)

func parsUrls(stringUrls []string) ([]*url.URL, error) {
	urls := slicez.Map(stringUrls, func(u string) result.Result[*url.URL] {
	    url
		return result.From(url.Parse(u))
	})
	return result.Unwrap(urls)
}

func main() {
	stringUrls := []string{
		"https://example.com",
		"https://github.com",
		"bad\n url",
	}
	urls, err := parsUrls(stringUrls)
	fmt.Println("URLs", urls)
	// URLs [https://example.com https://github.com]
	
	fmt.Println("Error", err)
    // Error parse "bad\n url": net/url: invalid control character in URL
}

```

## Content

Henry contain tree main packages. `slicez`, `chanz` and `mapz`


Functions in `slicez`
* Clone
* Compact
* CompactFunc
* Compare
* CompareFunc
* Complement
* ComplementBy
* Concat
* Contains
* ContainsFunc
* Cut
* CutFunc
* Difference
* DifferenceBy
* Drop
* DropRight
* DropRightWhile
* DropWhile
* Each
* Equal
* EqualFunc
* Every
* EveryFunc
* Filter
* Find
* FindLast
* FlatMap
* Flatten
* Fold
* FoldRight
* GroupBy
* Head
* Index
* IndexFunc
* Intersection
* IntersectionBy
* Join
* KeyBy
* Last
* LastIndex
* LastIndexFunc
* Map
* Max
* Min
* None
* NoneFunc
* Nth
* Partition
* Reject
* Reverse
* Sample
* Search
* Shuffle
* Some
* SomeFunc
* Sort
* SortFunc
* Tail
* Take
* TakeRight
* TakeRightWhile
* TakeWhile
* Union
* UnionBy
* Uniq
* UniqBy
* Unzip
* Unzip2
* Zip
* Zip2


Functions in `mapz`
* Clear
* Clone
* Copy
* DeleteFunc
* DeleteValue
* Equal
* EqualFunc
* Keys
* Merge
* Remap
* Values

Functions in `chanz`

* Collect
* CollectUntil
* Compact
* Compact1
* CompactN
* CompactUntil
* Concat
* Concat1
* ConcatN
* ConcatUntil
* Drop
* Drop1
* DropAll
* DropN
* DropUntil
* DropWhile
* DropWhile1
* DropWhileN
* DropWhileUntil
* EveryDone
* FanOut
* FanOut1
* FanOutN
* FanOutUntil
* Filter
* Filter1
* FilterN
* FilterUntil
* Flatten
* Flatten1
* FlattenN
* FlattenUntil
* Generate
* Generate1
* GenerateN
* GenerateUntil
* Map
* Map1
* MapN
* MapUntil
* Merge
* Merge1
* MergeN
* MergeUntil
* Partition
* Partition1
* PartitionN
* PartitionUntil
* Peek
* Peek1
* PeekN
* PeekUntil
* Readers
* SomeDone
* Take
* Take1
* TakeN
* TakeUntil
* TakeWhile
* TakeWhile1
* TakeWhileN
* TakeWhileUntil
* Unzip
* Unzip1
* UnzipN
* UnzipUntil
* Writers
* Zip
* Zip1
* ZipN
* ZipUntil



## Slicez

The `slicez` package contains generic utility functions and algorithms for slices

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

## Mapz
The `mapz` package contains generic utility functions and algorithms for maps

### Clear
Deletes every entry in a map

```go 
m := map[int]int{1:1, 2:2}
mapz.Clear(m)
// map[int]int{}
```

### Clone
Creates a clone of a map
```go 
m := map[int]int{1:1, 2:2}
mapz.Clone(m)
// map[int]int{1:1, 2:2}
```

### Copy
Copies one map into another

```go
src := map[int]int{1:1, 2:2}
dst := map[int]int{1:0, 3:3}
mapz.Copy(dst, src)
//  map[int]int{1:1, 2:2, 3:3}
```

### DeleteFunc
Will remove all entries from a map where the del function returns true

```go
m := map[int]int{1:1, 2:2, 3:3}
mapz.DeleteFunc(m, func(k, v int) bool { return k == 2 })
//  map[int]int{1:1, 3:3}
```

### DeleteValue
Deletes a value and the associated key from a map
```go
m := map[int]int{1:1, 2:800, 3:3}
mapz.DeleteValue(m, 800)
//  map[int]int{1:1, 3:3}
```

### Equal
Returns true if a map i equal

```go 
m1 := map[int]int{1:1, 3:3}
m2 := map[int]int{1:1, 3:3}
mapz.Equal(m1, m2)
// true
```


### EqualFunc
Returns true if a map i equal using the equality function to test it
```go 
m1 := map[int]int{1:1, 3:3}
m2 := map[int]int{1:1, 3:6}
mapz.EqualFunc(m1, m2, func(a, b int) bool {return a % 3 == b % 3})
// true
```

### Keys
Returns a slice of all keys in the map

```go 
m := map[int]int{1:1, 3:3}
mapz.Keys(m)
// []int{1,3}
```


### Merge
Merges multiple maps into one map
```go
m1 := map[int]int{1:1, 3:3, 4:800}
m2 := map[int]int{2:2, 4:4}
mapz.Merge(m1, m2)
// map[int]int{1:1, 2:2, 3:3, 4:4}
```

### Remap
Remaps a map in terms of its keys and values

```go 
m := map[int]int{2:2, 4:4}
mapz.Remap(m, func(k, v int) (k2, v2 string){
    return fmt.Sprint(k), fmt.Sprint(v*2)
})
// map[string]string{"2":"4", "4":"8"}
```

### Values
Returns a slice of all values in the map
```go 
m := map[int]int{2:6, 4:12}
mapz.Values(m)
// []int{6,12}
```


## Chanz

The `chanz` package contains generic utility functions and algorithms for channels


There are often 4 versions of the same function in the chanz package. They are simply shorthands for common configurations.
Examples of this is, `Map`, `Map1`, `MapN`, `MapUntil`.

* `Map` returns a channel of size `0` and will read from the input chan until it is closed
* `Map1` returns a channel of size `1` and will read from the input chan until it is closed
* `MapN` returns a channel of size `N` and will read from the input chan until it is closed
* `MapUntil` returns a channel of size `N` and will read from the input chan until it is closed or until the input `done` channel is closed 

### SomeDone
Takes N channels as input and returns one channel. If any of the input channels is closed, the output channel is closed. This 
is used for control structure.

```go 
done1 := make(chan, interface{})
done2 := make(chan, interface{})

done := chanz.SomeDone(done1, done2)

go func(){
   time.Sleep(time.Second)
    close(done1)
    time.Sleep(time.Second)
    close(done2)
}

<- done // will read in 1 secound
```



### EveryDone
Takes N channels as input and returns one channel. When all input channels is closed, the output channel will be closed.  This
is used for control structure.

```go 
done1 := make(chan, interface{})
done2 := make(chan, interface{})

done := chanz.SomeDone(done1, done2)

go func(){
    time.Sleep(time.Second)
    close(done1)
    time.Sleep(time.Second)
    close(done2)
}

<- done // will read in 2 seconds
```


### Collect, CollectUntil
Will collect all read items into a slice and return it

```go 
in := chanz.Generate(1,2,3,4,5)
chanz.Collect(in)
// []int{1,2,3,4,5}
```


### Compact, Compact1, CompactN, CompactUntil
Will remove consecutive duplicates from the channel

```go 
in := chanz.Generate(1,1,3,2,2,5,1)
w := chanz.Compact(in)
chanz.Collect(w)
// []int{1,3,2,5,1}
```




### Concat, Concat1, ConcatN, ConcatUntil
Will concatenate channels

```go 
in1 := chanz.Generate(1,2,3)
in2 := chanz.Generate(4,5,6)
w := chanz.Concat(in1, in2)
chanz.Collect(w)
// []int{1,2,3,4,5,6}
```

### Drop, Drop1, DropN, DropUntil
Drops the first N entries of the channel

```go 
in := chanz.Generate(1,2,3,4,5,6)
w := chanz.Drop(in, 2)
chanz.Collect(w)
// []int{3, 4, 5, 6}
```


### DropWhile, DropWhile1, DropWhileN, DropWhileUntil
Drops the first entries of the channel until function returns true

```go 
in := chanz.Generate(1,2,3,4,5,6,1)
w := chanz.DropWhile(in, func(i int) bool { return i < 3})
chanz.Collect(w)
// []int{3, 4, 5, 6, 1}
```


### DropAll
Drops all elements until closed

```go 
in := chanz.Generate(1,2,3,4,5,6)
w := chanz.DropAll(in, false)
chanz.Collect(w)
// []int{}
```

### DropBuffer
Drops all elements until closed

```go 
in := chanz.Generate1(1,2,3,4,5,6)
w := chanz.DropBuffer(in, false)
chanz.Collect(w)
// []int{2,3,4,5,6}
```

### FanOut, FanOut1, FanOutN, FanOutUntil
Takes an input channel and fans it out to multiple output channels
```go
in := chanz.Generate(1,2,3,4,5,6)
chans := chanz.FanOut(in, 2)
go chanz.Collect(chans[0])
chanz.Collect(chans[1])
// []int{1,2,3,4,5,6}
// []int{1,2,3,4,5,6}
```

### Filter, Filter1, FilterN, FilterUntil
Filters the items read onto the output chan

```go 
in := chanz.Generate(1,2,3,4,5,6)
even := chanz.Filter(in, func(i int) bool { return i % 2 == 0})
chanz.Collect(even)
// []int{2,4,6}
```

### Flatten, Flatten1, FlattenN, FlattenUntil
Flattens a channel that produces slices

```go 
in := chanz.Generate([]int{1,2,3}, []int{4,5,6})
w := chanz.Flatten(in)
chanz.Collect(w)
// []int{1,2,3,4,5,6}
```



### Generate, Generate1, GenerateN, GenerateUntil
Takes elements, creates a channel and writes the elements to it
```go
w := chanz.Generate(1,2,3,4)
chanz.Collect(w)
// []int{1,2,3,4}
```


### Map, Map1, MapN, MapUntil
Maps element from one channel to another

```go 
in := chanz.Generate(1,2,3,4)
w := chanz.Map(in, func(i int) string { return fmt.Sprint(i) })
chanz.Collect(w)
// []string{"1","2","3","4"}
```


### Merge, Merge1, MergeN, MergeUntil
Merge will take N chans and merge them onto one channel (in a non-particular order)

```go 
in1 := chanz.Generate(1,2,3)
in2 := chanz.Generate(4,5,6)
w := chanz.Merge(in1, in2)
chanz.Collect(w)
// []int{4,1,5,6,2,3}
```



### Partition, Partition1, PartitionN, PartitionUntil
Partition a channel into to two channels
```go 
in := chanz.Generate(1,2,3,4,5,6)
even, odd := chanz.Partition(in, func(i int) bool { return i % 2 == 0})
go chanz.Collect(even)
chanz.Collect(odd)
// []int{2,4,6}
// []int{1,3,5}
```


### Peek, Peek1, PeekN, PeekUntil
Will produce a channel that runs a function for each item

```go 
in := chanz.Generate(1,2,3,4,5,6)
in := chanz.Peek(in, func(i int){ fmt.Print(i)})
chanz.Collect(in)
// 123456
// []int{1,2,3,4,5,6}
```


### Take, Take1, TakeN, TakeUntil
Will take the first N items from the channel

```go 
in := chanz.Generate(1,2,3,4,5,6)
w := chanz.Take(in, 2)
chanz.Collect(w)
// []int{1,2}
```

### TakeWhile, TakeWhile1, TakeWhileN, TakeWhileUntil

Will take the first items from the channel until the predicate function returns false

```go 
in := chanz.Generate(1,2,3,4,5,6)
w := chanz.TakeWhile(in, func(i int) bool{ return i < 3})
chanz.Collect(w)
// []int{1,2}
```

### Unzip, Unzip1, UnzipN, UnzipUntil
Takes one chan and unzips it into two

```go 
in := chanz.Generate(-2, -1, 1, 2)
possitive, value := chanz.Unzip(in, func(i int) (bool, int) {
    return i > 0, Math.Abs(i)
})
go chanz.Collect(possitive)
chanz.Collect(value)
// []bool{false, false, true, true}
// []int{2,1,1,2}
```

### Zip, Zip1, ZipN, ZipUntil
Takes two channels and zips them into one channel 

```go 
in1 := chanz.Generate(1,2,3)
in2 := chanz.Generate("a","b","c")
x := chanz.Unzip(in1, in2, func(i int, s string) string {
    return fmt.Sprint(i,s)
})
chanz.Collect(w)
// []string{"1a", "2b", "3c"}
```


### Readers
Takes a slice of channels and returns a slice casted to read channels

### Writers
Takes a slice of channels and returns a slice casted to write channels



