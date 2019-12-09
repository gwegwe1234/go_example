package main

import "fmt"

func rptc(n *int) {
	*n = 77
}

func vptc(n int) {
	n = 77
}

func main() {
	// 포인터 값 전달
	// 함수, 메소드 호출 시에 매개변수 값을 복사 전달 -> 함수, 메소드 내에서는 원본 값 변경 불가능
	// 원본 값 변경 위해서 포인터 전달 -> ...? 이게 옳은걸까?
	// 특히 크기가 큰 배열인 경우 값 복사시에 시스템 부담 -> 포인터 전달로 해결(슬라이스, 맵 참조 전달)

	// 예제1

	var a int = 10
	var b int = 10

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println()

	rptc(&a)
	vptc(b)

	fmt.Println(a)
	fmt.Println(b)

	/*
		10
		10

		77
		10
	*/
}
