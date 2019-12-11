package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

// 인터페이스는 모든걸 다받아서, 이렇게 짜는사람도 많다.
func printValue(s interface{}) {
	fmt.Println("ex1 : ", s)
}

func main() {
	// 인터페이스 활용(빈 인터페이스)
	// 함수내에서 어떠한 타입이라도 유연하게 매개변수로 받을 수 있다.(만능) -> 모든 타입 지정 가능

	dog := Dog{"Poll", 10}
	cat := Cat{"Cat", 5}

	printValue(dog)
	printValue(cat)
	printValue(3)
	printValue("abcd")
}
