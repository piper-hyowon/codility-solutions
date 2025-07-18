package solution

import "sort"

func Solution(A []int) int {
	n := len(A)
	sort.Ints(A)

	maxProduct1 := A[n-1] * A[n-2] * A[n-3]
	maxProduct2 := A[0] * A[1] * A[n-1]

	if maxProduct1 > maxProduct2 {
		return maxProduct1
	}
	return maxProduct2
}
