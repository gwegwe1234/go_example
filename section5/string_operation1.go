package main

import "fmt"

func main()  {
	// 문자열 연산
	// 추출, 비교, 조합(결합)

	// 예제1 (추출)
	var str1 string = "Golang"
	var str2 string = "World"

	fmt.Println("ex1 : ", str1[0:2], str1[0])
	fmt.Println("ex1 : ", str2[3:], str2[0])
	fmt.Println("ex1 : ", str2[:4])
	fmt.Println("ex1 : ", str1[1:3])
	/*
		ex1 :  Go 71
		ex1 :  ld 87
		ex1 :  Worl
		ex1 :  ol
	 */
}
