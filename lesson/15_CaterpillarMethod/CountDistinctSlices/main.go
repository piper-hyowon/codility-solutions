// Task Score 70% Correctness 100% Performance 40%

package solution

import "fmt"

func Solution(M int, A []int) int {
	if len(A) == 1 {
		return 1
	}

	const (
		MAX = 1_000_000_000
	)

	cnt := 0
	// 애벌레
	tail, head := 0, 0

	for {
		if tail >= len(A) {
			break
		}
		if head >= len(A) {
			tail++
			head = tail
			continue
		}

		isUnique := true
		set := make(map[int]bool)
		for i := tail; i <= head; i++ {
			if set[A[i]] {
				isUnique = false
				break
			}
			set[A[i]] = true
		}

		// 다 유니크하면 애벌레 늘려
		if isUnique {
			fmt.Printf("(%d, %d) distinct!\n", tail, head)
			cnt++
			head++
		} else {
			fmt.Printf("(%d, %d) same!!\n", tail, head)
			// 같으면 꼬리 앞으로 이동해서길이 줄여
			tail++
			head = tail
		}

		if cnt >= MAX {
			return MAX
		}
	}

	return cnt
}

// (6,[1,5,1,5,2])
