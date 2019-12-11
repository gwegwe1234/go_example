package main

import (
	"fmt"
)

type Car struct { // 첫 글자 대문자
	name    string "차량명"
	color   string "색상"
	company string "제조사"
	detail  spec   "상세"
}

type spec struct {
	length int "전장"
	height int "전고"
	width  int "전축"
}

func main() {
	// 중첩구조체
	car1 := Car{
		"520d",
		"silver",
		"bmw",
		spec{4000, 1000, 2000},
	}

	fmt.Println(car1.name)
	fmt.Println(car1.color)
	fmt.Println(car1.company)
	fmt.Printf("%#v", car1.detail)

	fmt.Println(car1.detail.height)
}
