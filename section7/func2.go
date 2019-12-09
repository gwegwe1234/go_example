package main

import "fmt"

func sum2(i int, f func(int, int)) {
	f(i, 10)
}

func add(a int, b int) {
	fmt.Println("ex1 : ", a+b)
}

func multi_value(i int) {
	i = i * 10
}

func multi_reference(i *int) {
	*i = *i * 10
}

func main() {
	// 함수(콜백) 전달, 참조전달 (call by reference), 값 전달(call by value)
	sum2(10, add)

	a := 100
	multi_value(100)
	fmt.Println(a)

	multi_reference(&a)
	fmt.Println(a)
}
