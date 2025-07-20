// Task Score 55% Correctness 100% Performance 0%
package solution

func Solution(A []int) int {
	// 두 구간에서 리더면 어차피 A 에서 리더.
	var leader int
	half := len(A) / 2
	cnt := make(map[int]int)
	for _, v := range A {
		cnt[v]++
		if cnt[v] > half {
			leader = v
			break
		}
	}

	equiLeaders := 0
	for s := range len(A) - 1 {
		var first, second int
		for i, v := range A {
			if v == leader {
				if i <= s {
					first++
				} else {
					second++
				}
				if first > (s+1)/2 && second > (len(A)-s-1)/2 {
					equiLeaders++
					break
				}
			}
		}

	}

	return equiLeaders
}
