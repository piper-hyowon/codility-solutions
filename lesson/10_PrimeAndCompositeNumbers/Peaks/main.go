// Task Score 100% Correctness 100% Performance 100%

package solution

import "fmt"

// import "os"

func Solution(A []int) int {
	// 꼭대기를 최소한 하나가지게 같은 크기로 나눠.
	// 최대 몇개 블락이 생길수있는지
	// 나눌수 없으면 = 꼭대기가 하나도 없으면 0
	peaks := make([]int, 0, len(A))
	for i := 1; i < len(A)-1; i++ {
		if A[i-1] < A[i] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}
	}
	fmt.Println(peaks)
	if len(peaks) == 0 {
		return 0
	}

	blocks := len(peaks)
	// 꼭대기가 하나라도 있으면 1은 무조건 가능, 2까지만 확인
	for blocks > 1 {
		fmt.Println(blocks)
		peakIdx := 0
		k := len(A) / blocks
		curBlocks := 0 // 현재 블락 수
		for i := 0; i < len(A)-k+1; i += k {
			// i 부터 i + k안에 꼭대기 하나라도 있는지 확인
			for j := peakIdx; j < len(peaks); j++ {
				if peaks[peakIdx] >= i && peaks[peakIdx] < i+k {
					curBlocks++
					peakIdx++ // 이미 사용
					fmt.Printf("%d ~ %d has peak\n", i, i+k)
					break
				}
				fmt.Printf("%d ~ %d no peak\n", i, i+k)
			}
		}
		if curBlocks == blocks {
			return blocks
		}
		blocks--
	}

	return 1
}


package solution

import "sort"

func Solution(A []int) int {
	// 꼭대기를 최소한 하나가지게 같은 크기로 나눠.
	// 최대 몇개 블락이 생길수있는지
	// 나눌수 없으면 = 꼭대기가 하나도 없으면 0
	peaks := make([]int, 0, len(A))
	for i := 1; i < len(A)-1; i++ {
		if A[i-1] < A[i] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}
	}
	if len(peaks) == 0 {
		return 0
	}
    // len(A) 약수
    factors := make([]int, 0)
    for i := 1; i*i <= len(A); i++ {
        if len(A) % i == 0 {
            factors = append(factors, i)
            if i != len(A)/i {  // 제곱수가 아닐 때만
            factors = append(factors, len(A)/i)
            }
        }
    }
    sort.Slice(factors, func (i, j int) bool {
        return factors[i] > factors[j]
    })

	// 꼭대기가 하나라도 있으면 1은 무조건 가능, 2까지만 확인
	for _, blocks := range factors {
		peakIdx := 0
		k := len(A) / blocks
		curBlocks := 0 // 현재 블락 수
		for i := 0; i < len(A)-k+1; i += k {
			// i 부터 i + k안에 꼭대기 하나라도 있는지 확인
			for peakIdx < len(peaks) {
                if peaks[peakIdx] >= i + k {
                    break
                }
				if peaks[peakIdx] >= i && peaks[peakIdx] < i+k {
					curBlocks++
					peakIdx++ // 이미 사용
					break
				}
                peakIdx++
			}
		}
		if curBlocks == blocks {
			return blocks
		}
	}

	return 1
}
