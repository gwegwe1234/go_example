package main

import (
	"fmt"
	"math"
)

// f의 i 제곱을 구하는 함수

func Power(f float64, i float64) (float64, error) {
	if f == 0 {
		return 0, fmt.Errorf("0은 사용 불가능합니다")
	}

	return math.Pow(f, i), nil
}

func main() {

	// 예제 1
	if f, err := Power(7, 3); err != nil {
		fmt.Printf("Error Message : %s\n", err.Error())
	} else {
		fmt.Println("ex1 : ", f)
	}

	// 예제 2
	if f, err := Power(0, 3); err != nil {
		fmt.Printf("Error Message : %s\n", err)
	} else {
		fmt.Println("ex1 : ", f)
	}
}
