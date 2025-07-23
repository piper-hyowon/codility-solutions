package solution

import "math"

// 오름차순 배열! 서로 다른 절댓값 수 구하기
func Solution(A []int) int {
	set := make(map[int]bool)
	for _, v := range A {
		abs := int(math.Abs(float64(v)))
		if !set[abs] {
			set[abs] = true
		}
	}

	return len(set)
}
