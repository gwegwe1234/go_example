package main

import "fmt"

func main() {
	// 논리 연산자
	fmt.Println("ex1 : ", true && true)
	fmt.Println("ex1 : ", true && false)
	fmt.Println("ex1 : ", false && false)
	fmt.Println("ex1 : ", true || true)
	fmt.Println("ex1 : ", true || false)
	fmt.Println("ex1 : ", false || false)

	// 비교 연산자
	num1 := 15
	num2 := 37

	fmt.Println("ex2 : ", num1 < num2)
}
