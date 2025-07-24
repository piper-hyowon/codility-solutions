// Task Score 100% Correctness 100% Performance 100%

package solution

func Solution(A []int) int {
	set := make(map[int]bool)
	uniqueIdx := make(map[int]int) // 값: 인덱스
	for i, v := range A {
		if !set[v] {
			set[v] = true
			uniqueIdx[v] = i
		} else {
			// 또 등장. 유니크 X
			delete(uniqueIdx, v)
		}
	}
	min := len(A)
	result := 0
	for num, idx := range uniqueIdx {
		if idx < min {
			min = idx
			result = num
		}
	}

	if min < len(A) {
		return result
	}
	return -1
}
