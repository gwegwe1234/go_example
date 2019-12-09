package main

import "fmt"

func main() {

	cnt := increaseCnt()
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
	// 1, 2, 3, 4, 5
	// 지역변수가 캡쳐되면서 전역변수처럼 쓰인다
}

func increaseCnt() func() int {
	n := 0 // 지역변수(캡쳐됨)
	return func() int {
		n += 1
		return n
	}
}
