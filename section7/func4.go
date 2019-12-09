package main

import "fmt"

// 명시적으로 return 값의 이름을 정해줌
func multiply(x int, y int) (r1 int, r2 int) {
	r1 = x * 10
	r2 = y * 20
	return r1, r2
}

func main() {
	a, b := multiply(10, 5)

	fmt.Println(a, b)
}
