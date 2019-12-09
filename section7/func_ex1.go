package main

import "fmt"

func multiply(n ...int) int {
	tot := 1
	for _, value := range n {
		tot *= value
	}
	return tot
}

func sum(n ...int) int {
	tot := 1
	for _, value := range n {
		tot += value
	}

	return tot
}

func prtWord(msg ...string) {
	for _, value := range msg {
		fmt.Println("ex2 : ", value)
	}
}

func main() {
	// 함수 고급
	// 가변 인자 실습 (매개 변수 개수가 동적으로 변할 때 - 정해져 있지 않음)

	x := multiply(5, 6, 7, 8, 9, 10)
	y := sum(5, 6, 7, 8, 9)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println()

	prtWord("I", "am", "a", "boy")

	//
	a := []int{5, 6, 7, 8, 9}
	z := multiply(a...) // 슬라이스 내부의 인덱스로 알아서 접근해준다.
	n := sum(a...)
	fmt.Println(z, n)
}
