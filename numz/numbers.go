package numz

import (
	"github.com/modfin/henry/compare"
	"github.com/modfin/henry/slicez"
	"math"
	"sort"
)

// Min returns the minimum of the number supplied
func Min[N compare.Number](a ...N) N {
	return slicez.Min(a...)
}

// Max returns the maximum of the number supplied
func Max[N compare.Number](a ...N) N {
	return slicez.Max(a...)
}

// Range returns the range of the number supplied
func Range[N compare.Number](a ...N) N {
	return Max(a...) - Min(a...)
}

// Sum returns the sum of the number supplied
func Sum[N compare.Number](a ...N) N {
	var zero N
	return slicez.Fold(a, func(acc N, val N) N {
		return acc + val
	}, zero)
}

// VPow returns a vector containing the result of each element of "vector" to the power och "pow"
func VPow[N compare.Number](vector []N, pow N) []N {
	return slicez.Map(vector, func(a N) N {
		return N(math.Pow(float64(a), float64(pow)))
	})
}

// VMul will return a vector containing elements x and y multiplied such that x[i]*y[i] = returned[i]
func VMul[N compare.Number](x []N, y []N) []N {
	l := Min(len(x), len(y))
	x, y = x[:l], y[:l]
	return slicez.Zip(x, y, func(a, b N) N {
		return a * b
	})
}

// VAdd will return a vector containing elements x and y added such that x[i]+y[i] = returned[i]
func VAdd[N compare.Number](x []N, y []N) []N {
	l := Min(len(x), len(y))
	x, y = x[:l], y[:l]
	return slicez.Zip(x, y, func(a, b N) N {
		return a + b
	})
}

// VSub will return a vector containing elements y subtracted from x such that x[i]-y[i] = returned[i]
func VSub[N compare.Number](x []N, y []N) []N {
	l := Min(len(y), len(y))
	y, y = y[:l], y[:l]
	return slicez.Zip(x, y, func(a, b N) N {
		return a - b
	})
}

// VDot will return the dot product of two vectors
func VDot[N compare.Number](x []N, y []N) N {
	return Sum(VMul(x, y)...)
}

// Mean will return the mean of a vector
func Mean[N compare.Number](vector ...N) float64 {
	if len(vector) == 0 {
		var zero N
		return float64(zero)
	}
	return float64(Sum(vector...)) / float64(len(vector))
}

// MAD will return the Mean Absolute Deviation of a vector
func MAD[N compare.Number](vector ...N) float64 {
	mean := Mean(vector...)
	count := float64(len(vector))
	return slicez.Fold(vector, func(accumulator float64, val N) float64 {
		return accumulator + math.Abs(float64(val)-mean)/count
	}, 0.0)
}

// Var will return Variance of a sample
func Var[N compare.Number](samples ...N) float64 {
	avg := Mean(samples...)
	partial := slicez.Map(samples, func(x N) float64 {
		return math.Pow(float64(x)-avg, 2)
	})
	return Sum(partial...) / float64(len(partial)-1)
}

// StdDev will return the sample standard deviation
func StdDev[N compare.Number](samples ...N) float64 {
	return math.Sqrt(Var(samples...))
}

// StdErr will return the sample standard error
func StdErr[N compare.Number](n ...N) float64 {
	return StdDev(n...) / math.Sqrt(float64(len(n)))

}

// SNR will return the Signal noise ratio of the vector
func SNR[N compare.Number](sample ...N) float64 {
	return Mean(sample...) / StdDev(sample...)
}

// ZScore will return the z-score of vector
func ZScore[N compare.Number](x N, pop []N) float64 {
	return (float64(x) - Mean(pop...)) / StdDev(pop...)
}

// Skew will return the skew of a sample
func Skew[N compare.Number](sample ...N) float64 {
	count := float64(len(sample))
	mean := Mean(sample...)
	sd := StdDev(sample...)
	d := (count - 1) * math.Pow(sd, 3)

	return slicez.Fold(sample, func(accumulator float64, val N) float64 {
		return accumulator + math.Pow(float64(val)-mean, 3)/d
	}, 0)
}

// Corr returns the correlation between two vectors
func Corr[N compare.Number](x []N, y []N) float64 {
	l := Min(len(x), len(y))
	x, y = x[:l], y[:l]

	xm := Mean(x...)
	ym := Mean(y...)

	dx := slicez.Map(x, func(a N) float64 {
		return float64(a) - xm
	})
	dy := slicez.Map(y, func(a N) float64 {
		return float64(a) - ym
	})

	t := Sum(slicez.Zip(dx, dy, func(a float64, b float64) float64 {
		return a * b
	})...)

	n1 := Sum(VPow(dx, 2)...)
	n2 := Sum(VPow(dy, 2)...)

	return t / (math.Sqrt(n1 * n2))
}

// Cov returns the co-variance between two vectors
func Cov[N compare.Number](x []N, y []N) float64 {

	l := Min(len(x), len(y))
	x, y = x[:l], y[:l]

	xm := Mean(x...)
	ym := Mean(y...)

	dx := slicez.Map(x, func(a N) float64 {
		return float64(a) - xm
	})
	dy := slicez.Map(y, func(a N) float64 {
		return float64(a) - ym
	})

	t := Sum(slicez.Zip(dx, dy, func(a float64, b float64) float64 {
		return a * b
	})...)
	return t / float64(l)
}

// R2 returns the r^2 between two vectors
func R2[N compare.Number](x []N, y []N) float64 {
	return math.Pow(Corr(x, y), 2)
}

// LinReg returns the liniar regression of two verctors such that y = slope*x + intercept
func LinReg[N compare.Number](x []N, y []N) (intercept, slope float64) {
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

// FTest returns the F-Test of two vectors
func FTest[N compare.Number](x []N, y []N) float64 {
	return Var(x...) / Var(y...)
}

// Median returns the median of a vector
func Median[N compare.Number](vector ...N) float64 {
	l := len(vector)

	inter := slicez.SortBy(vector, func(i, j N) bool {
		return i < j
	})
	inter = slicez.Drop(inter, l/2-1)
	inter = slicez.DropRight(inter, l/2-1)

	if len(inter) == 2 {
		return Mean(inter...)
	}
	return float64(inter[len(inter)/2])
}

type modecount[N compare.Number] struct {
	val N
	c   int
}

// Mode returns the mode of a vector
func Mode[N compare.Number](vector ...N) N {
	return Modes(vector...)[0]
}

// Modes return the modes of all numbers in the vector
func Modes[N compare.Number](vector ...N) []N {
	m := map[N]int{}
	for _, n := range vector {
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

	inter := slicez.TakeWhile(counts, func(a modecount[N]) bool {
		return a.c == max
	})
	inter = slicez.SortBy(inter, func(a, b modecount[N]) bool {
		return a.val < b.val
	})

	return slicez.Map(
		inter,
		func(a modecount[N]) N {
			return a.val
		},
	)
}

// GCD returns the Greatest common divisor of a vector
func GCD[I compare.Integer](vector ...I) (gcd I) {
	if len(vector) == 0 {
		return gcd
	}
	gcd = vector[0]
	for _, b := range vector[1:] {
		for b != 0 {
			t := b
			b = gcd % b
			gcd = t
		}
	}
	return gcd
}

// LCM return the Least Common Multiple of a vector
func LCM[I compare.Integer](a, b I, vector ...I) I {
	result := a * b / GCD(a, b)

	for i := 0; i < len(vector); i++ {
		result = LCM(result, vector[i])
	}
	return result
}

// Percentile return "score" percentile of a vector
func Percentile[N compare.Number](score N, vector ...N) float64 {
	count := len(slicez.Filter(vector, func(n N) bool { return n < score }))
	return float64(count) / float64(len(vector))
}

// BitOR returns the bit wise OR between elements in a vector
func BitOR[I compare.Integer](vector []I) (i I) {
	if len(vector) == 0 {
		return i
	}
	if len(vector) == 1 {
		return vector[0]
	}
	return slicez.Fold(vector[1:], func(accumulator I, val I) I {
		return accumulator | val
	}, vector[0])
}

// BitAND returns the bit wise AND between elements in a vector
func BitAND[I compare.Integer](a []I) (i I) {
	if len(a) == 0 {
		return i
	}
	if len(a) == 1 {
		return a[0]
	}
	return slicez.Fold(a[1:], func(accumulator I, val I) I {
		return accumulator & val
	}, a[0])
}

// BitXOR returns the bit wise XOR between elements in a vector
func BitXOR[I compare.Integer](a []I) (i I) {
	if len(a) == 0 {
		return i
	}
	if len(a) == 1 {
		return a[0]
	}
	return slicez.Fold(a[1:], func(accumulator I, val I) I {
		return accumulator ^ val
	}, a[0])
}
