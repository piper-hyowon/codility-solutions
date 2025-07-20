// Task Score 100% Correctness 100% Performance 100%

package solution

// A[0] A[1] ... A[N-1]
// 상류 -> 하류

// B는 방향
// 0이면 상류로 가는중
// 1이면 하류로 흐름

// 만약 두 물고기가 반대방향으로 가고 있고 두 불고기 사이에 물고기가 ㅎ나ㅏ도 없다면,
// 언젠가 만난다. 만나면 한마리만 살아남음. 큰 물고기가 작은거 먹어
// 크기는 A[i]
// 속도는 다 같아서 같은 방향이면 절대 안만나.
// 살안마는 물고기 수를 구해라!!!

// (상류0) 0   1   2   3   4 (하류1)
//
//        <-  ->  <-  <-  <-

func Solution(A []int, B []int) int {
	stack := make([]int, 0, len(A))
	alive := 0

	for i := 0; i < len(A); i++ {
		if B[i] == 1 {
			// 하류로 가는 물고기
			stack = append(stack, i)
		} else {
			// 상류로 가는 물고기 등장
			for len(stack) > 0 {
				downstream := stack[len(stack)-1]

				if A[i] > A[downstream] {
					// 상류 물고기 승
					stack = stack[:len(stack)-1] // 하류 물고기 스택에서 제거
					// 상류 물고기는 계속 싸움
				} else {
					break // 루프 중단
				}
			}

			// 스택에 더 싸울 물고기 없음
			if len(stack) == 0 {
				alive++
			}
			// 중간에 졌으면 죽음
		}
	}

	alive += len(stack)

	return alive
}
