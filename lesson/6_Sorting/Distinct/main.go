// Task Score 100% Correctness 100% Performance 100%
package solution

func Solution(A []int) int {
	set := make(map[int]bool)
	for _, v := range A {
		if !set[v] {
			set[v] = true
		}
	}

	return len(set)
}
