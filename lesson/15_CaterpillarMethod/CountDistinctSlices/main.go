// Task Score 70% Correctness 100% Performance 40%

package solution

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
			cnt++
			head++
		} else {
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

// Task Score 100% Correctness 100% Performance 100%
// 슬라이스 중복 매번 확인 X
package solution

func Solution(M int, A []int) int {
	if len(A) == 1 {
		return 1
	}

	const  MAX = 1_000_000_000
	cnt := 0
	// 애벌레
	tail, head := 0, 0
    set := make(map[int]bool) // 재사용! 꼬리 앞으로가면 delete

	for {
		if tail >= len(A) {
			break
		}
		if head >= len(A) {
            delete(set, A[tail])
			tail++
			continue
		}

		// 다 유니크하면 애벌레 늘려
		if !set[A[head]] {
            set[A[head]] = true
            cnt += (head - tail + 1)
			head++
		} else {
            delete(set, A[tail])
			// 같으면 꼬리 앞으로 이동해서길이 줄여
			tail++
		}

		if cnt >= MAX {
			return MAX
		}
	}

	return cnt
}
