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

// Task Score 100% Correctness 100% Performance 100%
package solution
import "sort"

func Solution(A []int) int {
    n := len(A)
    if n < 2 {
        return 0
    }
    
    starts, ends := make([]int, n), make([]int, n)
    
    for i, radius := range A {
        starts[i], ends[i] = i-radius, i+radius
    }
    
    sort.Ints(starts)
    sort.Ints(ends)
    
    intersections := 0
    for i, j := 0, 0; i < n; i++ {
        // j를 적절한 위치로 이동
        for ; j < n && ends[j] < starts[i]; j++ {
            // 빈 body - j만 증가
        }
        
        if intersections += i - j; intersections > 10_000_000 {
            return -1
        }
    }
    
    return intersections
}