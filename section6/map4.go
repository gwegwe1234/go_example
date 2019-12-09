package main

import "fmt"

func main() {
	// 맵
	// 맵 조회 할 경우에 주의 할 점

	// 예제 1
	map1 := map[string]int{
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	value1, ok1 := map1["lemon"]
	value2 := map1["kiwi"]
	value3, ok := map1["kiwi"]

	fmt.Println("ex1 : ", value1, ok1)
	fmt.Println("ex1 : ", value2)
	fmt.Println("ex1 : ", value3, ok) // 두번째 리턴값으로 키 존재여부를 확인

	/*
		ex1 :  0 true //  레몬 키위 둘다 0이라 구분이 안가니까 ok라는 변수같은거로 true false 구분해주
		ex1 :  0
		ex1 :  0 false
	*/

	// 예제 2

	if value, ok := map1["kiwi"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("Kiwi is not exist")
	}

	// Kiwi is not exist

	if value, ok := map1["banana"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("banana is not exist")
	}

	// 115

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("Kiwi is not exist")
	}
	// Kiwi is not exist
}
