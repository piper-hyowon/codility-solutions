package solution

import "strconv"

func Solution(N int) int {
	binary := strconv.FormatInt(int64(N), 2)

	maxGap := 0
	currentGap := 0
	runes := []rune(binary)
	for i := 1; i < len(runes); i++ {
		if runes[i] == '0' {
			currentGap++
		} else {
			if currentGap > maxGap {
				maxGap = currentGap
			}
			currentGap = 0
		}

	}
	return maxGap
}
