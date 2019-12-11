package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

// 직원 안에 임원이 포함된거다. 즉 Employee 가 더 위에 단계
type Executives struct {
	Employee     // is a 관계. 즉 임원 is a 직원~
	specialBonus float64
}

func main() {

	// 구조체 임베디드 패턴
	// 다른 관점으로 메소드를 재 사용하는 장점 제공
	// 상속을 허용하지 않는 Go 언어에서 메소드 상속 활용을 위한 패턴

	// 예제1. 직원
	ep1 := Employee{"kim", 3000000, 150000}
	ep2 := Employee{"park", 5000000, 5000000}

	// 임원
	ex1 := Executives{
		Employee{"lee", 50000000, 10000000},
		1000000,
	}

	fmt.Println("ex1 : ", int(ep1.Calculate()))
	fmt.Println("ex1 : ", int(ep2.Calculate()))

	// Employee 구조체 통해서 임원이 메소드 호출하는거
	fmt.Println("ex2 : ", int(ex1.Calculate()+ex1.specialBonus))
}
