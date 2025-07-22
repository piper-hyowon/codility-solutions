// Task Score 66% Correctness 100% Performance 28%
package solution

func Solution(A []int) int {
	if len(A) <= 2 {
		return 0
	}

	// 일단 꼭대기를 찾아
	peaks := make([]int, 0, len(A)/2)
	for i := 1; i < len(A)-1; i++ {
		if A[i-1] < A[i] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}
	}
	if len(peaks) <= 1 {
		return len(peaks)
	}

	k := len(peaks)
	idx, cnt := 0, 1
	for idx < len(peaks)-1 && k > 0 {
		nextIdx := idx + 1
		for nextIdx < len(peaks) {
			if peaks[idx]+k <= peaks[nextIdx] {
				cnt++
				idx, nextIdx = nextIdx, nextIdx+1
				if cnt == k {
					return k
				}
			} else {
				nextIdx++
			}

		}

		// 거리 확보 안되면 k 줄여보기
		k--
		idx, cnt = 0, 1
	}

	return 0
}
