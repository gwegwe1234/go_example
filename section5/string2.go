package main

import "fmt"

func main() {
	// 문자열 표현
	// 문자열 : UTF-8 인코딩 (유니코드 문자 집합) -> 바이트 수 주의

	// 예제1
	var str1 string = "Golang"
	var str2 string = "World"
	var str3 string = "고프로그래밍"

	fmt.Println("ex1 : ", str1[0])
	fmt.Println("ex1 : ", str1[1])
	fmt.Println("ex1 : ", str1[2])
	fmt.Println("ex1 : ", str1[3])
	fmt.Println("ex1 : ", str1[4])
	fmt.Println("ex1 : ", str1[5])
	fmt.Println("ex1 : ", str2[0])
	fmt.Println("ex1 : ", str2[1])
	fmt.Println("ex1 : ", str2[2])
	fmt.Println("ex1 : ", str2[3])
	fmt.Println("ex1 : ", str2[4])
	fmt.Println("ex1 : ", str3[0])
	fmt.Println("ex1 : ", str3[1])
	fmt.Println("ex1 : ", str3[2])
	fmt.Println("ex1 : ", str3[3])
	fmt.Println("ex1 : ", str3[4])
	fmt.Println("ex1 : ", str3[5])

	// 예제 2
	fmt.Printf("ex1 : %c %c %c %c %c %c", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Println()

	conStr := []rune(str3) // 한글일 경우 이렇게
	fmt.Printf("ex1 : %c %c %c %c %c %c", conStr[0], conStr[1], conStr[2], conStr[3], conStr[4], conStr[5])
	fmt.Println()

	// 예제 3
	for i, char := range str1 {
		fmt.Printf("ex3 : %c(%d)\t", char, i)
	}
	fmt.Println()

	for i, char := range str2 {
		fmt.Printf("ex4 : %c(%d)\t", char, i)
	}
}
