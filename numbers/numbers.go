package numbers

import (
	"constraints"
	"github.com/crholm/henry"
	"github.com/crholm/henry/pipe"
	"math"
	"sort"
)

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

func VPow[N Numbers](a []N, pow N) []N {
	return henry.Map(a, func(_ int, a N) N {
		return N(math.Pow(float64(a), float64(pow)))
	})
}

func VMul[N Numbers](x []N, y []N) []N {
	l := Min(len(x), len(y))
	x, y = x[:l], y[:l]
	return henry.Zip(x, y, func(a, b N) N {
		return a * b
	})
}
func VAdd[N Numbers](x []N, y []N) []N {
	l := Min(len(x), len(y))
	x, y = x[:l], y[:l]
	return henry.Zip(x, y, func(a, b N) N {
		return a + b
	})
}
func VSub[N Numbers](x []N, y []N) []N {
	l := Min(len(y), len(y))
	y, y = y[:l], y[:l]
	return henry.Zip(x, y, func(a, b N) N {
		return a - b
	})
}

func VDot[N Numbers](x []N, y []N) N {
	return Sum(VMul(x, y)...)
}

func Mean[N Numbers](a ...N) float64 {
	if len(a) == 0 {
		var zero N
		return float64(zero)
	}
	return float64(Sum(a...)) / float64(len(a))
}

// MeanAbsDev - Mean Absolute Deviation
func MeanAbsDev[N Numbers](n ...N) float64 {
	mean := Mean(n...)
	count := float64(len(n))
	return henry.FoldLeft(n, func(_ int, accumulator float64, val N) float64 {
		return accumulator + math.Abs(float64(val)-mean)/count
	}, 0.0)
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
func StdErr[N Numbers](n ...N) float64 {
	return StdDev(n...) / math.Sqrt(float64(len(n)))

}

// SNR Signal noise ratio
func SNR[N Numbers](x ...N) float64 {
	return Mean(x...) / StdDev(x...)
}

func ZScore[N Numbers](x N, pop []N) float64 {
	return (float64(x) - Mean(pop...)) / StdDev(pop...)
}

func Skew[N Numbers](n ...N) float64 {
	count := float64(len(n))
	mean := Mean(n...)
	sd := StdDev(n...)
	d := (count - 1) * math.Pow(sd, 3)

	return henry.FoldLeft(n, func(_ int, accumulator float64, val N) float64 {
		return accumulator + math.Pow(float64(val)-mean, 3)/d
	}, 0)
}

// Corr Correlation
func Corr[N Numbers](x []N, y []N) float64 {
	l := Min(len(x), len(y))
	x, y = x[:l], y[:l]

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

	n1 := Sum(VPow(dx, 2)...)
	n2 := Sum(VPow(dy, 2)...)

	return t / (math.Sqrt(n1 * n2))
}

// Cov Covariance
func Cov[N Numbers](x []N, y []N) float64 {

	l := Min(len(x), len(y))
	x, y = x[:l], y[:l]

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
	return t / float64(l)
}

func R2[N Numbers](x []N, y []N) float64 {
	return math.Pow(Corr(x, y), 2)
}

func LinReg[N Numbers](x []N, y []N) (intercept, slope float64) {
	l := Min(len(x), len(y))
	x, y = x[:l], y[:l]

	sum_x := float64(Sum(x...))
	sum_x2 := float64(Sum(VPow(x, 2)...))
	sum_y := float64(Sum(y...))
	sum_xy := float64(Sum(VMul(x, y)...))
	n := float64(l)

	slope = (sum_y*sum_x2 - sum_x*sum_xy) / (n*sum_x2 - math.Pow(sum_x, 2))
	intercept = (n*sum_xy - sum_x*sum_y) / (n*sum_x2 - math.Pow(sum_x, 2))
	return intercept, slope
}

func FTest[N Numbers](x []N, y []N) float64 {
	return Variance(x...) / Variance(y...)
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

// GCD Greatest common divisor
func GCD[I constraints.Integer](ints ...I) (gcd I) {
	if len(ints) == 0 {
		return gcd
	}
	gcd = ints[0]
	for _, b := range ints[1:] {
		for b != 0 {
			t := b
			b = gcd % b
			gcd = t
		}
	}
	return gcd
}

// LCM Least Common Multiple
func LCM[I constraints.Integer](a, b I, integers ...I) I {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

func Percentile[N Numbers](score N, n ...N) float64 {
	count := pipe.Of(n).Filter(func(_ int, n N) bool { return n < score }).Count()
	return float64(count) / float64(len(n))
}

func BitOR[I constraints.Integer](a []I) (i I) {
	if len(a) == 0 {
		return i
	}
	if len(a) == 1 {
		return a[0]
	}
	return henry.FoldLeft(a[1:], func(_ int, accumulator I, val I) I {
		return accumulator | val
	}, a[0])
}

func BitAND[I constraints.Integer](a []I) (i I) {
	if len(a) == 0 {
		return i
	}
	if len(a) == 1 {
		return a[0]
	}
	return henry.FoldLeft(a[1:], func(_ int, accumulator I, val I) I {
		return accumulator & val
	}, a[0])
}

func BitXOR[I constraints.Integer](a []I) (i I) {
	if len(a) == 0 {
		return i
	}
	if len(a) == 1 {
		return a[0]
	}
	return henry.FoldLeft(a[1:], func(_ int, accumulator I, val I) I {
		return accumulator ^ val
	}, a[0])
}
