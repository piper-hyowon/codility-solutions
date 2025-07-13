// Task Score 100% Correctness 100% Performance 100%
package solution

import (
	"sort"
)

func Solution(A []int) int {
	sort.Ints(A)

	for len(A) >= 2 {
		if A[0] == A[1] {
			A = A[2:]
		} else {
			break
		}
	}

	return A[0]
}
