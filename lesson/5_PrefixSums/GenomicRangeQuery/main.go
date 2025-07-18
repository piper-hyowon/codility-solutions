// Detected time complexity: O(N * M)
// Task Score 62% Correctness 100% Performance 0%
package solution

func Solution(S string, P []int, Q []int) []int {
	factor := make(map[byte]int)
	factor['A'] = 1
	factor['C'] = 2
	factor['G'] = 3
	factor['T'] = 4
	result := make([]int, 0, len(P))
	for i := range P {
		min := 5
		for j := P[i]; j <= Q[i]; j++ {
			if factor[S[j]] < min {
				min = factor[S[j]]
			}
		}
		result = append(result, min)
	}

	return result
}

// Task Score 100% Correctness 100% Performance 100%
// "구간"에서 가장 작은 값 찾기 => 누적합! 
package solution

// 효율적인 솔루션 - O(N + M)
func Solution(S string, P []int, Q []int) []int {
    n := len(S)
    
    // 각 문자의 누적 개수를 저장할 2D 배열
    // prefixCount[i][j] = 0부터 i-1까지 문자 j의 개수
    // j: 0=A, 1=C, 2=G, 3=T
    prefixCount := make([][]int, n+1)
    for i := range prefixCount {
        prefixCount[i] = make([]int, 4)
    }
    
    // 누적합 계산
    for i := 0; i < n; i++ {
        // 이전 값 복사
        for j := 0; j < 4; j++ {
            prefixCount[i+1][j] = prefixCount[i][j]
        }
        
        // 현재 문자 카운트 증가
        switch S[i] {
        case 'A':
            prefixCount[i+1][0]++
        case 'C':
            prefixCount[i+1][1]++
        case 'G':
            prefixCount[i+1][2]++
        case 'T':
            prefixCount[i+1][3]++
        }
    }
    
    // 쿼리 처리
    result := make([]int, len(P))
    
    for i := range P {
        start, end := P[i], Q[i]
        
        // 구간 [start, end]에서 각 문자의 개수
        // A가 있는지 확인 (impact factor 1)
        if prefixCount[end+1][0] - prefixCount[start][0] > 0 {
            result[i] = 1
        // C가 있는지 확인 (impact factor 2)
        } else if prefixCount[end+1][1] - prefixCount[start][1] > 0 {
            result[i] = 2
        // G가 있는지 확인 (impact factor 3)
        } else if prefixCount[end+1][2] - prefixCount[start][2] > 0 {
            result[i] = 3
        // T만 있음 (impact factor 4)
        } else {
            result[i] = 4
        }
    }
    
    return result
}