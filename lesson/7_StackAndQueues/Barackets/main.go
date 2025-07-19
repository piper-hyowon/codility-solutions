// Task Score 100% Correctness 100% Performance 100
package solution

func Solution(S string) int {
	if len(S) == 0 {
		return 1
	}

	stack := make([]rune, 0)
	for _, v := range []rune(S) {
		switch v {
		case '(', '[', '{':
			stack = append(stack, v)
		case ')', '}', ']':
			if len(stack) < 1 {
				return 0
			}
			pair := string(stack[len(stack)-1]) + string(v)
			if pair != "()" && pair != "{}" && pair != "[]" {
				return 0
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) > 0 {
		return 0
	}
	return 1
}
