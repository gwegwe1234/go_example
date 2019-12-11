package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

//bite 메소드
func (d Dog) bite() {
	fmt.Println(d.name, "bites!")
}

// 동물의 행동 인터페이스 선언
type Behavior interface {
	bite()
}

func main() {
	// 인터페이스 구현 예제
	// 예제 1
	dog1 := Dog{"poll", 10}

	var inter1 Behavior
	inter1 = dog1
	inter1.bite()

	// 예제 2
	dog2 := Dog{"marry", 12}
	inter2 := Behavior(dog2)

	inter2.bite()

	// 예제 3
	inters := []Behavior{dog1, dog2}

	for idx, _ := range inters {
		inters[idx].bite()
	}

	// 값 형태로 실행 (인터페이스)
	for _, val := range inters {
		val.bite()
	}
}
