package main

import "math"

func main() {
	// 예제1 오버플로우 에러
	var n1 uint8 = math.MaxUint8 + 1
}
