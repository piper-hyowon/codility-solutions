// Task Score 100% Correctness 100% Performance

package solution

func Solution(X int, A []int) int {
	// Implement your solution here
	// A의 원소가 1 ~ X가 다 등장해야함. 가장 작은 i 를 구하기.
	set := make(map[int]bool)
	// set 의 크기가 X가 될 때의 i 를 리턴

	for i := 0; i < len(A); i++ {
		if A[i] <= X && !set[A[i]] {
			set[A[i]] = true
		}

		if len(set) == X {
			return i
		}
	}

	return -1
}
