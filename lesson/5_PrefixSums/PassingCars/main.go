// Task Score 60% Correctness 100% Performance 20%
package solution

func Solution(A []int) int {
	if len(A) == 1 {
		return 0
	}
	passing := 0
	for P := 0; P < len(A)-1; P++ {
		for Q := P + 1; Q < len(A); Q++ {
			if A[P] == 0 && A[Q] == 1 {
				passing++
				if passing > 1000000000 {
					return -1
				}
			}
		}
	}

	return passing
}

// 누적합!!
// Task Score 100% Correctness 100% Performance 100%
package solution

func Solution(A []int) int {
    eastCar := 0
    passingCars := 0
    for _, v := range A {
        if v == 0 {
            eastCar++
        } else {
            // 앞에 있는 모든 eastCars와 passing pair를 만듦
            passingCars += eastCar
            if passingCars > 1000000000 {
                return -1
            }
        }
    }

    return passingCars
}
