// Task Score 93% Correctness 100% Performance 87%
package solution

import "sort"

func Solution(A []int) int {
	// 중심 - 반지름 ~ 중심 + 반지름
	// 이 영역이 겹치는지!
	path := make([][]int, 0, len(A))
	for i, v := range A {
		path = append(path, []int{i - v, i + v})
	}
	// 시작점 빠른 순
	sort.Slice(path, func(i, j int) bool {
		return path[i][0] < path[j][0]
	})

	cnt := 0
	for i, v := range path {
		tmp := 0
		for j := i + 1; j < len(path); j++ {
			// v[1] 랑 비교
			if path[j][0] <= v[1] {
				tmp++
			} else {
				break
			}
		}
		cnt += tmp
	}

	if cnt > 10_000_000 {
		return -1
	}

	return cnt
}
