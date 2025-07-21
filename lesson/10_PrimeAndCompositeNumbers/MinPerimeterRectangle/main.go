// Task Score 100% Correctness 100% Performance 100%

package solution

import "math/big"

// 넓이를 줄테니까 가로세로 내가 정해서 최소 둘레를 알아내라는거지?
// A * B 를 줄테니 2(A + B) 의 최솟값을 구해라!
// A, B는 N의 약수, A, B차이가 가장 작아야해..
func Solution(N int) int {
	// 소수면 답은 하나밖에 없음
	if big.NewInt(int64(N)).ProbablyPrime(1) {
		return 2 * (1 + N)
	}

	divide := 1
	end := N
	for divide <= end {
		if N%divide == 0 {
			end = N / divide
		}
		divide++
	}
	return 2 * (end + N/end)
}
