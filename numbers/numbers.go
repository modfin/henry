package numbers

import (
	"constraints"
	"github.com/crholm/henry"
	"github.com/crholm/henry/pipe"
	"math"
	"sort"
)

type Numbers interface {
	constraints.Integer | constraints.Float
}

func Min[N Numbers](a ...N) N {
	if len(a) == 0 {
		panic("no min of nothing")
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func Max[N Numbers](a ...N) N {
	if len(a) == 0 {
		panic("no max of nothing")
	}
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

func Range[N Numbers](a ...N) N {
	return Max(a...) - Min(a...)
}

func Sum[N Numbers](a ...N) N {
	var zero N
	return henry.FoldLeft(a, func(_ int, acc N, val N) N {
		return acc + val
	}, zero)
}

func Pow[N Numbers](a []N, pow N) []N {
	return henry.Map(a, func(_ int, a N) N {
		return N(math.Pow(float64(a), float64(pow)))
	})
}

func Mean[N Numbers](a ...N) float64 {
	if len(a) == 0 {
		var zero N
		return float64(zero)
	}
	return float64(Sum(a...)) / float64(len(a))
}

func Variance[N Numbers](samples ...N) float64 {
	avg := Mean(samples...)
	partial := henry.Map(samples, func(_ int, x N) float64 {
		return math.Pow(float64(x)-avg, 2)
	})
	return Sum(partial...) / float64(len(partial)-1)
}

func StdDev[N Numbers](samples ...N) float64 {
	return math.Sqrt(Variance(samples...))
}

func Correlation[N Numbers](x []N, y []N) float64 {
	xm := Mean(x...)
	ym := Mean(y...)

	dx := henry.Map(x, func(_ int, a N) float64 {
		return float64(a) - xm
	})
	dy := henry.Map(y, func(_ int, a N) float64 {
		return float64(a) - ym
	})

	t := Sum(henry.Zip(dx, dy, func(a float64, b float64) float64 {
		return a * b
	})...)

	n1 := Sum(Pow(dx, 2)...)
	n2 := Sum(Pow(dy, 2)...)

	return t / (math.Sqrt(n1 * n2))
}

func Median[N Numbers](n ...N) float64 {
	l := len(n)
	middle := pipe.Of(n).
		Sort(func(i, j N) bool {
			return i < j
		}).
		DropLeft(l/2 - 1).
		DropRight(l/2 - 1).
		Slice()
	if len(middle) == 2 {
		return Mean(middle...)
	}
	return float64(middle[len(middle)/2])
}

type modecount[N Numbers] struct {
	val N
	c   int
}

func Mode[N Numbers](nums ...N) N {
	return Modes(nums...)[0]
}

func Modes[N Numbers](nums ...N) []N {
	m := map[N]int{}
	for _, n := range nums {
		m[n] = m[n] + 1
	}

	var counts []modecount[N]
	for n, c := range m {
		counts = append(counts, modecount[N]{n, c})
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].c > counts[j].c
	})
	max := counts[0].c
	return henry.Map(
		pipe.Of(counts).
			TakeLeftWhile(func(i int, a modecount[N]) bool {
				return a.c == max
			}).
			Sort(func(a, b modecount[N]) bool {
				return a.val < b.val
			}).
			Slice(),
		func(i int, a modecount[N]) N {
			return a.val
		},
	)
}

func GCD[I constraints.Integer](a, b I) I {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM[I constraints.Integer](a, b I, integers ...I) I {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
