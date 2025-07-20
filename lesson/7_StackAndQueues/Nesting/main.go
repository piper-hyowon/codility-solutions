// Task Score 100% Correctness 100% Performance 100%

package solution

func Solution(S string) int {
	opened := 0

	for _, v := range S {
		if v == '(' {
			opened++
		} else {
			if opened >= 1 {
				opened--
			} else {
				return 0
			}
		}
	}
	if opened > 0 {
		return 0
	}

	return 1
}
