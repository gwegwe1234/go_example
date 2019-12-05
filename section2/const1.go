package main

import "fmt"

func main() {
	// 상수
	// const 사용 초기화, 한 번 선언 후 값 변경 금지
	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10
	// const d = getHeight()
	const e = 35.6
	const f = false

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	fmt.Println("e : ", e)
	fmt.Println("f : ", f)

}
