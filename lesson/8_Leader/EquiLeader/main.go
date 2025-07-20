// Task Score 55% Correctness 100% Performance 0% (시간복잡도: O(n²) )
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

// Task Score 100% Correctness 100% Performance 100%
package solution

func Solution(A []int) int {
	// 두 구간에서 리더면 어차피 A 에서 리더.
	var leader int
	half := len(A) / 2
	cnt := make(map[int]int)
	leaderCnt := 0
	for _, v := range A {
		cnt[v]++
		if cnt[v] > half {
			leader = v
			leaderCnt = cnt[v]
			// break 제거: 리더를 찾아도 총 개수를 파악 필요
		}
	}
	if leaderCnt == 0 {
		return 0  
	}

	equiLeaders := 0
    leadersInFirst, leadersInSecond := 0, leaderCnt
	for s := 0; s < len(A) - 1; s++ {
        if (A[s] == leader) {
            leadersInFirst++
            leadersInSecond--
        }
        if leadersInFirst > (s+1)/2 && leadersInSecond > (len(A)-s-1)/2 {
                equiLeaders++
		}
	}

	return equiLeaders
}


