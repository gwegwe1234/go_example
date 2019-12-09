package main

import (
	"fmt"
)

func main() {
	// 함수 고급
	// 익명함수(Anonymous Function)

	// 선언과 동시에 실행할 때 사용

	func() {
		fmt.Println("익명함수!")
	}()

	msg := "Hello Golang"

	func(m string) {
		fmt.Println(m)
	}(msg)

	func(x, y int) {
		fmt.Println(x + y)
	}(3, 4)

	r := func(x, y int) int {
		return x * y
	}(10, 100)

	fmt.Println(r)
}
