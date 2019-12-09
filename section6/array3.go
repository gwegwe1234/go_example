package main

import "fmt"

func main() {
	// 배열 복사
	// 값 복사 확인

	// 예제 1
	arr1 := [5]int{1, 10, 100, 1000, 10000}
	arr2 := arr1

	fmt.Println("ex1 : ", arr1, &arr1)
	fmt.Println("ex1 : ", arr2, &arr2)

	fmt.Printf("ex1 : %p %v\n", &arr1, arr1)
	fmt.Printf("ex1 : %p %v\n", &arr2, arr2)

	/*
		ex1 :  [1 10 100 1000 10000] &[1 10 100 1000 10000]
		ex1 :  [1 10 100 1000 10000] &[1 10 100 1000 10000]
		ex1 : 0xc0000a0000 [1 10 100 1000 10000]
		ex1 : 0xc0000a0030 [1 10 100 1000 10000]
	*/
}
