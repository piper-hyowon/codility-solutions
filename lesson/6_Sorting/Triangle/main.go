// Task Score 100% Correctness 100% Performance 100%

package solution

import "sort"

func Solution(A []int) int {
	if len(A) < 3 {
		return 0
	}
	sort.Ints(A)
	for i, v := range A {
		if v > 0 {
			A = A[i:]
			break
		}
	}
	for i := 0; i < len(A)-2; i++ {
		if A[i]+A[i+1] > A[i+2] {
			return 1
		}
	}
	return 0
}
