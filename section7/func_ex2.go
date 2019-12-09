package main

import "fmt"

func multiply(x, y int) (r int) {
	r = x * y
	return r
}

func sum(x, y int) (r int) {
	r = x + y
	return r
}

func main() {
	// 함수 고급
	// 함수를 변수에 할당

	// 예제 1 (슬라이스에 할당)
	f := []func(int, int) int{multiply, sum}

	a := f[0](10, 10)
	b := f[1](10, 10)
	fmt.Println(a, b) //......

	var f1 func(int, int) int = multiply
	f2 := sum

	fmt.Println(f1(10, 10))
	fmt.Println(f2(10, 10))

	// 맵에 할당
	m := map[string]func(int, int) int{
		"mul_func": multiply,
		"sum_func": sum,
	}

	fmt.Println(m["mul_func"](10, 10))
	fmt.Println(m["sum_func"](10, 10))
}
