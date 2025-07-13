// Task Score 44% Correctness 40% Performance 50%
package solution

import "sort"

// A에 없고, 0보다 큰 가장 작은 정수
func Solution(A []int) int {
	if len(A) == 1 {
		if A[0] == 1 {
			return 2
		} else {
			return 1
		}
	}

	sort.Ints(A)
	if A[0] <= 0 {
		return 1
	}

	for i := 0; i < len(A)-1; i++ {
		if A[i] >= 0 && A[i+1]-A[i] > 1 {
			return A[i] + 1
		}
	}

	return A[len(A)-1] + 1
}

// early return 조건... 정렬 후 가장 작은수가 1보다 크면 1리턴
// 가장 큰 수가 0이하면 다 음수라는거니까 이것도 1리턴
// Task Score 88% Correctness 80%\ Performance 100% 아놔 
package solution

import "sort"

// A에 없고, 0보다 큰 가장 작은 정수
func Solution(A []int) int {
    if len(A) == 0 {
        return 1
    }
	if len(A) == 1 {
		if A[0] == 1 {
			return 2
		} else {
			return 1
		}
	}

	sort.Ints(A)
	if A[0] > 1  || A[len(A) - 1] <= 0 {
		return 1
	}

	for i := 0; i < len(A)-1; i++ {
		if A[i] >= 0 && A[i+1]-A[i] > 1 {
			return A[i] + 1
		}
	}

	return A[len(A)-1] + 1
}

// 3트
// Task Score 100% Correctness 100% Performance 100%
package solution

import "sort"

func Solution(A []int) int {
    smallest := 1
    sort.Ints(A)
    for _, v := range A {
        if v <= 0 {
            continue
        }
        if smallest == v {
            smallest++
        } else if v > smallest {
            return smallest
        }
    }

    return smallest
}
