// Task Score 77% Correctness 100% Performance 50

package solution

// 언제 사고 언제 팔아야 최대 이익?
// 최대 이익 리턴, 최대 이익이 마이너스면 0 리턴
func Solution(A []int) int {
	if len(A) == 0 {
		return 0
	}
	buy, sell := 0, 1 // index
	profits := []int{}
	cheapest := 0 // index
	// buy 는 항상 sell 보다 작아야함(먼저 사야되니깐)
	for buy < len(A)-2 && sell < len(A)-1 {
		p := A[sell] - A[buy]
		// 수익이 마이너스면 차라리 판 날에 사고,
		// 다음날 팔아보자
		if p < 0 {
			buy, sell = sell, sell+1
			continue
		}
		// 더 비싸게 팔 수 있으면 갱신
		if A[sell+1] > A[sell] {
			sell++
			// profits = append(profits, p)
		} else if A[sell+1] > A[buy] {
			// 수익은 더 안나지만 산 가격보단 비싸면 다른 가격 또 알아보자
			sell++
		} else {
			// 산 가격보다 저렴하면 이때 살까? 지금 수익 킵해놓고! <- 이게 틀렸네.자꾸지나쳐.가장 저렴한걸 킵해두자
			profits = append(profits, p)
			if A[cheapest] > A[buy] {
				cheapest = buy
			}
			buy, sell = sell+1, sell+2
		}
	}

	max := 0
	if buy >= 0 && sell >= 0 && max < A[sell]-A[buy] {
		max = A[sell] - A[buy]
	}
	for _, v := range profits {
		if v > max {
			max = v
		}
	}

	return max
}
