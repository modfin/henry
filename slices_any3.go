package henry

func Zip[A any, B any, C any](aSlice []A, bSlice []B, zipper func(a A, b B) C) []C {
	var i = len(aSlice)
	var j = len(bSlice)
	if j < i {
		i = j
	}
	var cSlice []C
	for k, a := range aSlice {
		if k == j {
			break
		}
		b := bSlice[k]
		cSlice = append(cSlice, zipper(a, b))
	}
	return cSlice
}

func Unzip[A any, B any, C any](cSlice []C, unzipper func(c C) (a A, b B)) ([]A, []B) {
	var aSlice []A
	var bSlice []B
	for _, c := range cSlice {
		a, b := unzipper(c)
		aSlice = append(aSlice, a)
		bSlice = append(bSlice, b)
	}
	return aSlice, bSlice

}
