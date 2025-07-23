// Task Score 62% Correctness 100% Performance 25%

package solution

// 0 ~ N-1 초콜릿
// 먹고나면 포장지만 남아
// X 초콜릿 먹어 -> (X + M) modulo N 초콜릿먹어?
// 빈 포장지 만날때까지 먹어.
func Solution(N int, M int) int {
	eaten := make(map[int]bool)
	eaten[0] = true
	i := M % N
	for {
		if eaten[i%N] {
			return len(eaten)
		}
		eaten[i%N] = true
		if len(eaten) >= N {
			break
		}
		i += M
	}

	return len(eaten)
}
