package main

import "fmt"

func a() {
	defer end(start("b")) // defer 문은 처음 감싸는 함수만 적용. 즉 end는 적용되지만, start는 적용 안됨.
	fmt.Println("in a")
}

func start(msg string) string {
	fmt.Println("start : ", msg)
	return msg
}

func end(msg string) {
	fmt.Println("end : ", msg)
}

func main() {
	// 예제 1
	a()
	/*
		start :  b
		in a
		end :  b
	*/
}
