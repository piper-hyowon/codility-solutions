// Task Score 100% Correctness 100% Performance 100%

package solution

// 선분 안겹치는 원소들로만 집합을만들때 집합최대 길이.
func Solution(A []int, B []int) int {
	if len(A) == 0 {
		return 0
	}

	// 앞 선분 끝나는 점이랑
	// 뒤 선분 시작 점 비교
	end := 0
	start := 1
	cnt := 1
	for end < len(A)-1 && start < len(A) {
		// 안겹쳐
		if B[end] < A[start] {
			cnt++
			end, start = start, start+1
		} else {
			// 겹쳐
			start++
		}
	}

	return cnt
}
