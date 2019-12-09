package main

import "fmt"

func main()  {
	// Bool : 참과 거짓
	// 조건부 논리 연산자랑 주로 사용 : !, ||, &&
	// 암묵적 형변환 x

	// 예제 1
	var b1 bool = true
	var b2 bool = false
	b3 := true

	fmt.Println(b1, b2, b3)

	fmt.Println(b3 == b3)
	fmt.Println(b1 && b3)
	fmt.Println(!b2)
	fmt.Println(b2 || b3)
}
