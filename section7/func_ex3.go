package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func prtHello(n int) {
	if n == 0 {
		return
	}

	fmt.Println("ex2 : (", n, ") hi!")
	prtHello(n - 1)
}

func main() {
	// 함수 고급
	// 재귀 함수
	// 프로그래이 보기 쉽고, 코드 갼결, 오류 수정이 용이 : 장점
	// 코드이해하기 어렵다. 기억공간을 많이 사용. 무한루프 가능성이 있음

	// 예제 1
	x := fact(7)

	fmt.Println(x)

	prtHello(10)
}
