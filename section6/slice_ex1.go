package main

import "fmt"

func main() {
	// 슬라이스 추가 및 병합
	// 예제 1

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{8, 9, 10, 11, 12}
	s3 := []int{13, 14, 15, 16, 17}
	s1 = append(s1, 6, 7)
	s2 = append(s1, s2...)      // 슬라이스를 삽입 할 경우 ... 사용
	s3 = append(s2, s3[0:3]...) // 추출 후 병합

	fmt.Println("ex1 : ", s1)
	fmt.Println("ex1 : ", s2)
	fmt.Println("ex1 : ", s3)

	/*
		ex1 :  [1 2 3 4 5 6 7]
		ex1 :  [1 2 3 4 5 6 7 8 9 10 11 12]
		ex1 :  [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]
	*/

	// 예제 2
	s4 := make([]int, 0, 5)
	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
		fmt.Printf("ex2 -> len : %d, cap : %d, value: %v\n", len(s4), cap(s4), s4)
	}

	/*
		 	ex2 -> len : 1, cap : 5, value: [0]
			ex2 -> len : 2, cap : 5, value: [0 1]
			ex2 -> len : 3, cap : 5, value: [0 1 2]
			ex2 -> len : 4, cap : 5, value: [0 1 2 3]
			ex2 -> len : 5, cap : 5, value: [0 1 2 3 4]
			ex2 -> len : 6, cap : 10, value: [0 1 2 3 4 5] // 용량이 모자라면, append는 2배로 용량을 늘려버린다.
			ex2 -> len : 7, cap : 10, value: [0 1 2 3 4 5 6]
			ex2 -> len : 8, cap : 10, value: [0 1 2 3 4 5 6 7]
			ex2 -> len : 9, cap : 10, value: [0 1 2 3 4 5 6 7 8]
			ex2 -> len : 10, cap : 10, value: [0 1 2 3 4 5 6 7 8 9]
			ex2 -> len : 11, cap : 20, value: [0 1 2 3 4 5 6 7 8 9 10]
			ex2 -> len : 12, cap : 20, value: [0 1 2 3 4 5 6 7 8 9 10 11]
			ex2 -> len : 13, cap : 20, value: [0 1 2 3 4 5 6 7 8 9 10 11 12]
			ex2 -> len : 14, cap : 20, value: [0 1 2 3 4 5 6 7 8 9 10 11 12 13]
			ex2 -> len : 15, cap : 20, value: [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14]
	*/

}
