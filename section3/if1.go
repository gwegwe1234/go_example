package main

import "fmt"

func main()  {
	// 조건문
	// 반드시 Boolean 검사 1,0 이딴거 안됨
	// 소괄호 x

	var a int = 20
	b := 20

	// 예제 1
	if a >= 15 {
		fmt.Println("15 이상")
	}

	if b >= 25 {
		fmt.Println("25 이상")
	}

	// 애러 발생 1
	/*
	if b >= 25
	{
		// 이렇게 하면 에러발생
	}
	*/

	// 에러 발생 2
	/*
	if b >= 25
		fmt.Println("25이상")
		중괄호 생략 불가
	*/

	// 에러 발생 3
	/*
	if c:=1; c {
		// 에러발생. 1따위로 안된다.
	}
	*/

	if c := 40; c >= 35 {
		println("35 이상")
	}
}
