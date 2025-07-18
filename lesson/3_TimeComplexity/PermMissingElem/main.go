// 100
package solution

import "sort"

func Solution(A []int) int {
	n := len(A)
	if n == 0 {
		return 1
	}

	sort.Ints(A)

	// 첫 번째 원소가 1이 아닌 경우
	if A[0] != 1 {
		return 1
	}

	// 중간에 빠진 수 찾기
	for i := 0; i < n-1; i++ {
		if A[i]+1 != A[i+1] {
			return A[i] + 1
		}
	}

	// 마지막 수가 빠진 경우
	return n + 1
}
