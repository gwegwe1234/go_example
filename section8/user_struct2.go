package main

import "fmt"

type cnt int

func main() {
	// 기본 자료형 사용자 정의 타입
	// 예제 1

	a := cnt(15)
	fmt.Println("ex1 : ", a)

	// 예제 2

	var b cnt = 15
	fmt.Println("ex2 : ", b)

	// 예제 3 -> cnt형을 int에다가 넣는거 안된다
	//testConvertT(b)
	testConvertT(int(b)) // 형변환
	testConvertD(b)
}

func testConvertT(i int) {
	fmt.Println("ex3 : (Default Type) : ", i)
}

func testConvertD(i cnt) {
	fmt.Println("ex4 : (Custom Type) : ", i)
}
