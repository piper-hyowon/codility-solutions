// Task Score 66% Correctness 100% Performance 25%
package solution

import (
	"strconv"
)

func Solution(A []int) []int {
	result := make([]int, 0, len(A))
	memory := make(map[string]int) // 1 이면 약수아님. 2이면 약수
	counts := make(map[int][]int)  // 값 : [1,non-divisor 수]
	for i, a := range A {
		if len(counts[a]) >= 1 {
			result = append(result, counts[a][1])
			continue
		}
		nonDivisors := 0
		for j, b := range A {
			// 자신과의 비교는 패스
			if i == j {
				continue
			}
			//  b가 a 의 약수이려면 a이하여야함, 초과는 약수가될수없음
			if a < b {
				nonDivisors++
				continue
			}
			key := strconv.Itoa(a) + "," + strconv.Itoa(b)
			// 약수 아님
			if memory[key] == 1 {
				nonDivisors++
			} else if memory[key] == 2 {
				// 약수임 패스
				continue
			} else {
				// 저장된 정보 없음
				// 약수 임
				if a >= b && a%b == 0 {
					memory[key] = 2
				} else {
					nonDivisors++
					memory[key] = 1
				}
			}
		}

		result = append(result, nonDivisors)
		counts[a] = []int{1, nonDivisors}
	}
	return result
}
