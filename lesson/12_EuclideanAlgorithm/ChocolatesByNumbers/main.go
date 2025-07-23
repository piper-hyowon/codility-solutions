// Task Score 100% Correctness 100% Performance 100%
package solution

// M을 몇 번 더해야 N의 배수가 되는지
// => 최소공배수 구하고 / M
// LCM(10, 4) = 20
// 0 4 8 12 16 20(0)
// 20이 되어서 인덱스가 다시 0이 되면 안먹음
// 그 전에 16까지 5번먹음! 20되려면 4를 5번 더해. 20/4가 필요
// 최소공배수 = A X B / 최대공약수
func Solution(N int, M int) int {
	return lcm(N, M, gcd(N, M)) / M
}

func lcm(a, b, gcd int) int {
	return a * b / gcd
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
