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

func (d Dog) bite() {
	fmt.Println(d.name, "Dog bite!")
}

func (d Dog) sound() {
	fmt.Println(d.name, "Dog bark!")
}

func (d Dog) run() {
	fmt.Println(d.name, "Dog is running!")
}

// 구조체 Cat 메소드 구현
func (c Cat) bite() {
	fmt.Println(c.name, "Cat bite!")
}

func (c Cat) sound() {
	fmt.Println(c.name, "Cat cries!")
}

func (c Cat) run() {
	fmt.Println(c.name, "Cat is running!")
}

// 동물의 행동 인터페이스 선언
type Behavior interface {
	bite()
	sound()
	run()
}

// 인터페이스를 파라미터 받는다.
func act(animal Behavior) {
	animal.bite()
	animal.sound()
	animal.run()
}

func main() {
	// 덕타이핑 : 구조체 및 변수의 값이나 타입은 상관하지 않고 오로지 구현된 메소드로만 판단하는 방식
	dog := Dog{"poll", 10}
	cat := Cat{"bob", 5}

	// 개 행동 실행
	act(dog)
	act(cat)
	// 고양이 행동 실행
}
