// 100 점 (for문)
package solution

func Solution(A []int, K int) []int {
	if len(A) == 1 || len(A) == 0 || K%len(A) == 0 {
		return A
	}

	result := make([]int, len(A), len(A))
	for i := 0; i < len(A); i++ {
		result[(K+i)%len(A)] = A[i]
	}

	return result
}

// 1트, K 음수 인거 생각 안 함 (20 몇 점)
package solution

func Solution(A []int, K int) []int {
	n := len(A)
	if n == 0 || n == 1 || K%n == 0 {
		return A
	}

	K = K % n
	// A[i]들이 다 A[i+K] 로 이동해야하는데
	// 기존: A[0:n]
	// 바꾼 후: A[K-1:n]..., A[0:K-1]...

	return append(append([]int{}, A[K-1:n]...), A[0:K-1]...)
}

// K 음수 고려 100점! 
func Solution(A []int, K int) []int {
    n := len(A)
    if n == 0 {
        return A
    }

    K = ((K % n) + n) % n
    if K == 0 {
        return A
    }

    return append(append([]int{}, A[n-K:]...), A[:n-K]...)
}

