package main

import (
	"fmt"
)

func main() {
	// 슬라이스 복사
	// copy(복사 대상, 원본)
	// make로 공간을 할당 후 복사 해야한다.
	// 복사 된 슬라이스 값 변경해도 원본에는 영향 없음

	// 예제1 (복사)
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1)
	copy(slice3, slice1)

	fmt.Println("ex1 : ", slice2)
	fmt.Println("ex1 : ", slice3)

	/*
		ex1 :  [1 2 3 4 5] // slice2가 공간이 5개라 5개만 복사됨
		ex1 :  []
	*/

	// 예제2
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)
	copy(b, a)

	b[0] = 7
	b[4] = 10

	fmt.Println(a, b)

	/*
		[1 2 3 4 5] [7 2 3 4 10]
	*/

	// 예제3
	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3]

	d[1] = 7

	fmt.Println(c, d)
	/*
		[1 7 3 4 5] [1 7 3] --> 추출을 해도 원본값은 바뀐다. 즉 참조하고 있다.
	*/

	// 예제 4
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7]

	fmt.Println(len(f), cap(f), f)

	/*
		5 7 [1 2 3 4 5] // 용량 7로 복사
	*/
}
