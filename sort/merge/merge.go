package merge

/*
EXAMPLE USAGE

func main() {
	// Example input
	input := []int64{5, 6, 3, 7, 4, 9, 8, 1, 11, 3, 4, 2, 9, 10, 7, 1}
	fmt.Println("Result V1", SortV1(input))
	fmt.Println("Result V2", SortV2(input))
}
*/

// SortV1 : Do a merge sort algorithm (or so i thought)
func SortV1(arr []int64) []int64 {
	var result []int64

	if len(arr) == 1 {
		return arr
	}

	l := arr[:len(arr)/2]
	r := arr[len(arr)/2:]

	lp := SortV1(l)
	rp := SortV1(r)

	result = mergeV1(lp, rp)

	return result
}

func mergeV1(l []int64, r []int64) []int64 {
	var result []int64

	for j := 0; j < len(l); j++ {
		for i := 0; i < len(r); i++ {
			if l[j] <= r[i] {
				l[j], r[i] = r[i], l[j]
			}
		}
	}

	r = SortV1(r)

	result = append(l, r...)
	return result
}

// SortV2 : Do a merge sort algorithm with correct & simpler moves
func SortV2(arr []int64) []int64 {
	var result []int64

	if len(arr) == 1 {
		return arr
	}

	l := arr[:len(arr)/2]
	r := arr[len(arr)/2:]

	lp := SortV2(l)
	rp := SortV2(r)

	result = mergeV2(lp, rp)

	return result
}

func mergeV2(l []int64, r []int64) []int64 {
	var result []int64
	// using pointer from each arr
	iL, iR := 0, 0

	for j := 0; j < len(append(l, r...)); j++ {
		if iR >= len(r) { // right depleted?
			result = append(result, l[iL]) // append current Left
			iL++                           // shift Left position
			continue
		} else if iL >= len(l) { // left depleted?
			result = append(result, r[iR]) // append current Right
			iR++                           // shift Right position
			continue
		}

		if l[iL] <= r[iR] { // left smaller?
			result = append(result, r[iR]) // append current Right
			iR++                           // shift Right position
		} else {
			result = append(result, l[iL]) // append current Left
			iL++                           // shift Left position
		}
	}
	return result
}
