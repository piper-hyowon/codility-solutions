// Task Score 100% Correctness 100% Performance 100%
package solution

import "sort"

func Solution(A []int) int {
	if len(A) <= 2 {
		return 0
	}

	sort.Ints(A)
	// 1 2 5 8 10 12
	// one + two > three
	cnt := 0
	one := 0

	for one <= len(A)-3 {
		for two := one + 1; two < len(A)-1; two++ {
			three := two + 1
			for three < len(A) {
				if A[one]+A[two] > A[three] {
					three++
				} else {
					break
				}
			}
			cnt += three - two - 1
		}
		one++
	}

	return cnt
}
