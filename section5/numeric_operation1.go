package main

import (
	"fmt"
	"math"
)

func main()  {
	// 숫자 연산 (산술, 비교)
	// 타입이 같아야 가능
	// 다른 타입끼리는 반드시 형 변환 후 연산
	// 형 변환 없을 경우 -> 에러 발생
	// +, -, *, %, /, <<, >>, ^

	// 예제1
	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(math.MaxFloat64)

	n5 := 10000
	n6 := int16(10000)
	n7 := uint8(100)

	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n6 + int16(n7))

	fmt.Println(n6 > int16(n7))

}