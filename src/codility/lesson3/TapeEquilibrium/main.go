// task score: 69 correctness: 100 Performance: 33
package solution

import "math"

func Solution(A []int) int {
	difference := make([]int, 0, len(A)-1)
	for p := 1; p < len(A); p++ {
		first := 0
		second := 0
		for i := 0; i < p; i++ {
			first += A[i]
		}
		for i := p; i < len(A); i++ {
			second += A[i]
		}

		difference = append(difference, int(math.Abs(float64(first)-float64(second))))
	}

	min := difference[0]
	for i := 1; i < len(difference); i++ {
		if difference[i] < min {
			min = difference[i]
		}
	}

	return min
}


// 슬라이딩 윈도우  task score: 100 correctness: 100 Performance: 100
package solution

import "math"

func Solution(A []int) int {
	difference := make([]int, 0, len(A)-1)
    var sum func (i, j int) int
    sum = func (i, j int) int {
        if i == j {
            return A[i]
        }

        if j - i == 1 {
            return A[i] + A[j]
        }
        return sum(i, j - 1) + A[j]
    }

    difference = append(difference, )

	for p := 1; p < len(A); p++ {
		difference = append(difference, int(math.Abs(float64(sum(0, p - 1)) - float64(sum(p, len(A) - 1)))))
	}

	min := difference[0]
	for i := 1; i < len(difference); i++ {
		if difference[i] < min {
			min = difference[i]
		}
	}

	return min
}



package solution

func Solution(A []int) int {
    leftSum := A[0]
    rightSum := 0
    for i := 1; i < len(A); i ++ {
        rightSum += A[i]
    }

    abs := func (a, b int) int {
        if a > b {
            return a - b
        } else {
            return b - a
        }
    }
    min := abs(leftSum, rightSum)

    for leftEnd := 1; leftEnd < len(A) - 1; leftEnd++ {
        leftSum += A[leftEnd]
        rightSum -= A[leftEnd]
        diff := abs(leftSum, rightSum)
        if diff < min {
            min = diff
        }
    }

    return min
}
