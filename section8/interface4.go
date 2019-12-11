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

// 구조체 Dog 메소드 구현
func (d Dog) run() {
	fmt.Println(d.name, "Dog is running!")
}

// 구조체 Cat 메소드 구현
func (c Cat) run() {
	fmt.Println(c.name, "Cat is running!")
}

func act(animal interface{ run() }) {
	animal.run()
}

func main() {
	// 익명 예제

	dog := Dog{"poll", 10}
	cat := Cat{"bob", 5}

	// 개 행동 실행
	act(dog)
	act(cat)
	// 고양이 행동 실행
}
