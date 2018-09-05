package merge

import (
	"fmt"
)

func main() {
	input := []int64{5, 6, 3, 7, 4, 9, 8, 1, 11, 3, 4, 2, 9, 10, 7, 1}
	fmt.Println("Result V1", SortV1(input))
}

// SortV1 : Do a merge sort algorithm (or so i thought)
func SortV1(arr []int64) []int64 {
	var splitted []int64

	if len(arr) == 1 {
		return arr
	}

	l := arr[:len(arr)/2]
	r := arr[len(arr)/2:]

	fmt.Println("before splitmerge", l, r)
	lp := SortV1(l)
	rp := SortV1(r)

	fmt.Println("after splitmerge", lp, rp)
	splitted = mergeV1(lp, rp)

	return splitted
}

func mergeV1(l []int64, r []int64) []int64 {

	var swapped []int64

	for j := 0; j < len(l); j++ {
		for i := 0; i < len(r); i++ {
			if l[j] <= r[i] {
				l[j], r[i] = r[i], l[j]
			}
		}
	}

	r = SortV1(r)

	swapped = append(l, r...)

	return swapped
}
