package main

import "fmt"

func main() {
	// 맵(map)
	// 맵 : 해시테이블, 딕셔너리(파이썬)
	// 레퍼런스 타입(참조 값 전달)
	// 비교연산자 사용 불가능
	// 특징 : 참조타입(key)로 사용 불가능, 값(Value)로는 모든 타입 사용 가능
	// make 함수 및 축약(리터럴)로 초기화 가능
	// 순서 없음

	// 예제 1
	var map1 map[string]int = make(map[string]int) // key : string value : int
	var map2 = make(map[string]int)
	map3 := make(map[string]int) // 주로 이렇게

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)
	fmt.Println()

	/*
		map[]
		map[]
		map[]
	*/

	// 예제 2
	map4 := map[string]int{} //JSON 형태
	map4["apple"] = 25
	map4["banana"] = 40
	map4["orange"] = 33

	map5 := map[string]int{
		"apple":  15,
		"banana": 40,
		"orange": 33,
	}

	fmt.Println(map4)
	fmt.Println(map5)
	fmt.Println()

	/*
		map[apple:25 banana:40 orange:33]
		map[apple:15 banana:40 orange:33]
	*/
	map6 := make(map[string]int, 10)
	map6["apple"] = 25
	map6["banana"] = 40
	map6["orange"] = 33

	fmt.Println(map6)
	fmt.Println(map6["orange"])

	/*
		map[apple:25 banana:40 orange:33]
		33
	*/
}
