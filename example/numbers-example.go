package main

import (
	"fmt"
	"github.com/crholm/go18exp/numberz"
	"github.com/crholm/go18exp/slicez"
)

func main() {

	ints := []int{1, 2, 3, 4, 5, 23, 12, 5, 231, 12, 12}
	fmt.Println("Integers")
	fmt.Println("Mean:", numberz.Mean(ints...))
	fmt.Println("Median:", numberz.Median(ints...))
	fmt.Println("Mode:", numberz.Mode(ints...))
	fmt.Println("Min:", numberz.Min(ints...))
	fmt.Println("Max:", numberz.Max(ints...))
	fmt.Println("Sum:", numberz.Sum(ints...))
	fmt.Println("x^2:", numberz.VPow(ints, 2))
	fmt.Println("Variance:", numberz.Var(ints...))  // Sample
	fmt.Println("StdDev:", numberz.StdDev(ints...)) // Sample
	fmt.Println()

	floats := slicez.Map(ints, numberz.MapFloat64[int])
	floats = slicez.Map(floats, func(a float64) float64 {
		return a + 0.5
	})
	fmt.Println("Floats")
	fmt.Println("Mean:", numberz.Mean(floats...))
	fmt.Println("Median:", numberz.Median(floats...))
	fmt.Println("Mode:", numberz.Mode(floats...))
	fmt.Println("Min:", numberz.Min(floats...))
	fmt.Println("Max:", numberz.Max(floats...))
	fmt.Println("Sum:", numberz.Sum(floats...))
	fmt.Println("x^2:", numberz.VPow(floats, 2))
	fmt.Println("Variance:", numberz.Var(floats...))
	fmt.Println("StdDev:", numberz.StdDev(floats...))

}
