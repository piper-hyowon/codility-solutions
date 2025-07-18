// Task Score 100% Correctness 100% Performance 100%

package solution

import "math"

func Solution(X int, Y int, D int) int {
	return int(math.Ceil(float64(Y-X) / float64(D)))
}
