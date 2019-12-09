package main

import "fmt"

func multiply(x int, y int) (int, int) {
	return x * 10, y * 10
}

func arrayMultiply(a, b, c, d, e int) (int, int, int, int, int) {
	return a, b * 2, c * 3, d * 4, e * 5
}

func main() {
	// 다중 값 반환
	// 예제 1
	a, b := multiply(10, 5)
	c, _ := multiply(10, 5)
	_, d := multiply(10, 5)

	fmt.Println(a, b, c, d)

	// 예제 2
	x1, x2, x3, x4, x5 := arrayMultiply(1, 2, 3, 4, 5)
	y1, y2, _, y4, _ := arrayMultiply(1, 2, 3, 4, 5)
	fmt.Println(x1, x2, x3, x4, x5)
	fmt.Println(y1, y2, y4)

}
