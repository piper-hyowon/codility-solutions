// Task Score 100% Correctness 100% Performance 100%
package solution

import "math"

// 평균이 최소가 되는 구간의 시작 인덱스 리턴
func Solution(A []int) int {
	minAvg := float64(math.MaxInt32)
	minIdx := 0

	// 구간합 구하기
	// 길이 2
	for i := range len(A) - 1 {
		avg := float64(A[i]+A[i+1]) / 2.0
		if avg < minAvg {
			minAvg = avg
			minIdx = i
		}
	}

	// 길이 3
	for i := 0; i < len(A)-2; i++ {
		avg := float64(A[i]+A[i+1]+A[i+2]) / 3.0
		if avg < minAvg {
			minAvg = avg
			minIdx = i
		}
	}

	return minIdx
}
