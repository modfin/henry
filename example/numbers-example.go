package main

import (
	"fmt"
	"github.com/crholm/henry"
	"github.com/crholm/henry/numbers"
)

func main() {

	ints := []int{1, 2, 3, 4, 5, 23, 12, 5, 231, 12, 12}
	fmt.Println("Integers")
	fmt.Println("Mean:", numbers.Mean(ints...))
	fmt.Println("Median:", numbers.Median(ints...))
	fmt.Println("Mode:", numbers.Mode(ints...))
	fmt.Println("Min:", numbers.Min(ints...))
	fmt.Println("Max:", numbers.Max(ints...))
	fmt.Println("Sum:", numbers.Sum(ints...))
	fmt.Println("x^2:", numbers.Pow(ints, 2))
	fmt.Println("Variance:", numbers.Variance(ints...)) // Sample
	fmt.Println("StdDev:", numbers.StdDev(ints...))     // Sample
	fmt.Println()

	floats := henry.Map(ints, numbers.MapFloat64[int])
	floats = henry.Map(floats, func(_ int, a float64) float64 {
		return a + 0.5
	})
	fmt.Println("Floats")
	fmt.Println("Mean:", numbers.Mean(floats...))
	fmt.Println("Median:", numbers.Median(floats...))
	fmt.Println("Mode:", numbers.Mode(floats...))
	fmt.Println("Min:", numbers.Min(floats...))
	fmt.Println("Max:", numbers.Max(floats...))
	fmt.Println("Sum:", numbers.Sum(floats...))
	fmt.Println("x^2:", numbers.Pow(floats, 2))
	fmt.Println("Variance:", numbers.Variance(floats...))
	fmt.Println("StdDev:", numbers.StdDev(floats...))

}
