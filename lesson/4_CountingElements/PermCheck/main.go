// Task Score 75% Correctness 83% Performance 66%
// sum으로만 하면 잘못된 점 -> 총합은 순열이 아니여도 같을 수 있음..!
package solution

func Solution(A []int) int {
	requiredSum := len(A) * (len(A) + 1) / 2
	sum := 0
	for _, v := range A {
		sum += v
	}
	if sum != requiredSum {
		return 0
	} else {
		return 1
	}
}

// 중복된 원소가 있는지도 검사. 있으면 빨리 0 리턴
// Task Score 100% Correctness 100% Performance 100%
package solution

func Solution(A []int) int {
    set := make(map[int]bool)
    requiredSum := len(A) * (len(A) + 1) / 2
    sum := 0
    for _, v := range A {
        if set[v] {
            return 0
        }
        sum += v
        set[v] = true
    }

    if requiredSum == sum {
        return 1
    }
    return 0
}
