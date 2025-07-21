// Task Score 100% Correctness 100% Performance 100%

package solution

// Kadane's Algorithm
func Solution(A []int) int {
	maxSum, curSum := A[0], A[0]
	for i := 1; i < len(A); i++ {
		// 이전 합이 음수면 버리고 새로 시작
		// 양수면 계속 더하기
		if curSum < 0 {
			curSum = A[i]
		} else {
			curSum += A[i]
		}
		if curSum > maxSum {
			maxSum = curSum
		}
	}

	return maxSum
}
