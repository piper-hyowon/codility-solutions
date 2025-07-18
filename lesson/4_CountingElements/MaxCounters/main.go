// Task Score 77% Correctness 100% Performance 60%
// 몇몇케이스에서 타임아웃..
package solution

func Solution(N int, A []int) []int {
	counters := make([]int, N, N)
	max := 0
	for _, X := range A {
		if X == N+1 {
			for i, _ := range counters {
				counters[i] = max
			}
		} else {
			counters[X-1]++
			if counters[X-1] > max {
				max = counters[X-1]
			}
		}
	}

	return counters
}


// Task Score 100% Correctness 100% Performance 100%
// Lazy Update
// max counter를 기록만 해두고, 필요할 때 업데이트
// 마지막에 아직 업데이트 안 된 카운터들 처리
package solution

func Solution(N int, A []int) []int {
	counters := make([]int, N, N)
	max := 0
	lastMaxCounter := 0
	for _, X := range A {
		if X == N+1 {
			lastMaxCounter = max
		} else {
			if counters[X-1] < lastMaxCounter {
				counters[X-1] = lastMaxCounter
			}

			counters[X-1]++
			if counters[X-1] > max {
				max = counters[X-1]
			}
		}
	}

	for i := range counters {
		if counters[i] < lastMaxCounter {
			counters[i] = lastMaxCounter
		}
	}

	return counters
}
