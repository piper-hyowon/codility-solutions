// Task Score 100% Correctness 100% Performance 100

package solution

func Solution(A []int) int {
	half := len(A) / 2
	cnt := make(map[int]int)
	for i, v := range A {
		cnt[v]++
		if cnt[v] > half {
			return i
		}
	}

	return -1
}
