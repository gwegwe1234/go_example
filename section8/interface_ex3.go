package main

import "fmt"

func checkType(arg interface{}) {
	// arg.(type)
	switch arg.(type) {
	case bool:
		fmt.Println("this is a bool", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("this is int", arg)
	case float64:
		fmt.Println("this is float", arg)
	case string:
		fmt.Println("this is string", arg)
	default:
		fmt.Println("this is nil", arg)
	}
}

func main() {
	// 실제 타입 검사 switch
	// 빈 인터페이스는 어떠한 자료형도 전달 받을 수 있으므로, 타입 체크를 통해 형 변환 후 사용 가능

	// 예제 1
	checkType(true)
	checkType(1)
	checkType(22.41)
	checkType(nil)
	checkType("Hello Golang")
}
