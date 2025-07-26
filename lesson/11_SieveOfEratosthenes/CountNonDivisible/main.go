// Task Score 55% Correctness 100% Performance 0%

package solution

import (
	"fmt"
	"strconv"
)

// import "math/big"

func Solution(A []int) []int {
	// big.ProbablyPrime(big.NewInt(n))
	// 각 원소마다 약수의 개수배열 리턴?
	result := make([]int, 0, len(A))
	memory := make(map[string]int) // 1 이면 약수아님. 2이면 약수
	for i, a := range A {
		fmt.Printf("%d (%d)\n", i, a)
		nonDivisors := 0
		for j, b := range A {
			fmt.Print(b)
			// 자신과의 비교는 패스
			if i == j {
				fmt.Println("self pass")
				continue
			}
			//  b가 a 의 약수이려면 a이하여야함, 초과는 약수가될수없음
			if a < b {
				nonDivisors++
				fmt.Println("non-divisors (a < b)")
				continue
			}
			key := strconv.Itoa(a) + "," + strconv.Itoa(b)
			// 약수 아님
			if memory[key] == 1 {
				nonDivisors++
				fmt.Println("non-divisors (in memory)")
			} else if memory[key] == 2 {
				// 약수임 패스
				fmt.Println("divisors (in memory)")
				continue
			} else {
				// 저장된 정보 없음
				// 약수 임
				if a >= b && a%b == 0 {
					memory[key] = 2
					fmt.Println("divisors (new memory)")
				} else {
					nonDivisors++
					memory[key] = 1
					fmt.Println("non-divisors (new memory)")
				}
			}
		}
		result = append(result, nonDivisors)
	}
	return result
}
