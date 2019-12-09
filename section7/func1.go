package main

import (
	"fmt"
	"strconv"
)

func helloGolang() {
	fmt.Println("ex1 : Hello Golang!")
}

func say_one(m string) {
	fmt.Println("ex2 : ", m)
}

func sum(x int, y int) int {
	return x + y
}

func main() {
	// 함수
	// 선언 : func 키워드로 선언
	// func 함수명(매개변수) (반환타입 or 반혼 값 변수명) : 반환 값 존재
	// func 함수명() (반환 타입 or 반환 값 변수명) : 매개변수 없음, 반환값 존재
	// func 함수명(매개변수) : 매개변수 존재, 반환값 없음
	// func 함수명() : 둘다 없음
	// 타 언어와 달리 반환값 다수 가능

	helloGolang()

	say_one("One!")

	num := sum(5, 5)
	fmt.Println(num)
	fmt.Println(strconv.Itoa(num))
}
