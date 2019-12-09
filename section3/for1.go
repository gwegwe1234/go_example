package main

import "fmt"

func main() {

	// 예제 1
	for i := 0; i < 5; i++ {
		fmt.Println("ex1: ", i)
	}

	// 에러 발생 case

	/*
		for i := 0; i < 5; i++
		{
			fmt.Println(중괄호 내리면 에러)
		}
	*/

	/*
		for i := 0; i < 5; i++
			fmt.Println("괄호 없으면 에러")
	*/

	// 예제 2 무한루프
	/*
		for {
			fmt.Println("ex2 : Hello, Golang!")
			fmt.Println("ex2 : Infinite loog")
		}
	*/

	// 예제 3 (Range 용법)
	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println("ex3: ", index, name)
	}
}
