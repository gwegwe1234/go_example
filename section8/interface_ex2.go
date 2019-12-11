package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 타입 변환 -> type Assertion
	// 실행(런타임) 시에는 인터페이스에 할당한 변수는 실제 타입으로 변환 후 사용해야 하는 경우
	// 인터페이스.(타입) -> 형변환
	// interfaceValue.(type)

	// 예제 1
	var a interface{} = 15
	b := a
	c := a.(int)
	//d := a.(float64) -> 이런거 안됨. int 형이라 int 로 변환만 된다. 런타임 에러

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(a))
	//fmt.Println(d)

	// 예제 2(저장된 실제 타입 검사)
	if v, ok := a.(int); ok {
		fmt.Println("ex2 : ", v)
	}

}
