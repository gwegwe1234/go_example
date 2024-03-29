package main

import "fmt"

func main() {
	// 짧은 선언
	// 반드시 함수 안에서만 사용(전역 x), 선언 후 할당 예외 발생
	// 주로 제한된 범위의 함수내에서 사용 할 경우 코드 가독성을 높일 수 있다. -> 추천

	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false

	// shortVar1 := 10 이러면 에러난다.

	fmt.Println("shortVar1 : ", shortVar1)
	fmt.Println("shortVar2 : ", shortVar2)
	fmt.Println("shortVar3 : ", shortVar3)

	if i := 10; i < 11 {
		fmt.Println("Short Variable If test")
	}
}
