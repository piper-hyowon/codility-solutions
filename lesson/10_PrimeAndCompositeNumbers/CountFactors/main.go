// Task Score 100% Correctness 100% Performance 100%

package solution

import "math/big"

func Solution(N int) int {
	if N == 1 {
		return 1
	}

	// 소수면 1이랑 자기 자신
	if big.NewInt(int64(N)).ProbablyPrime(1) {
		return 2
	}
	factors := make(map[int]bool)
	set := func(n int) {
		if !factors[n] {
			factors[n] = true
		}
	}
	end := N
	i := 2
	for i <= end {
		if N%i == 0 {
			set(i)
			set(N / i)
			end = N / i
		}
		i++
	}

	return len(factors) + 2 // 1, N 추가
}
