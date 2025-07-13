// Task Score 100% Correctness 100% Performance 100%
package solution

func Solution(A int, B int, K int) int {
	if A == 0 {
		return B/K + 1
	}

	return B/K - (A-1)/K
}
