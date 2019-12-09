package main

import "fmt"

func main() {
	// Go : 모호하거나, 애매한 문법 금지
	// 후치 연산자 허용 i++, 전치 연산자 비허용 ++i
	// 증감연산 반환 값 없음 sum := i+= 안됨
	// 포인터(Pointer 허용, 연산 비허용)
	// 주석 (//, /* */)

	// 예제 1
	sum, i := 0, 0
	for i <= 100 {
		// sum += i++ 이런거 안된다.
		sum += i
		i++
	}

	fmt.Println("ex1 : ", sum)
}
