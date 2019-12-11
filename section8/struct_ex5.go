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

// 메소드 오버라이딩 ..! 개신기하네
func (e Executives) Calculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

type Executives struct {
	Employee
	specialBonus float64
}

func main() {

	ep1 := Employee{"kim", 3000000, 150000}
	ep2 := Employee{"park", 5000000, 5000000}
	ex1 := Executives{
		Employee{"lee", 50000000, 10000000},
		1000000,
	}

	fmt.Println("ex1 : ", int(ep1.Calculate()))
	fmt.Println("ex1 : ", int(ep2.Calculate()))

	fmt.Println("ex2 : ", int(ex1.Calculate())) // 오버라이드 된다.. 신기신기
	fmt.Println("ex2 : ", int(ex1.Employee.Calculate()+ex1.specialBonus))
}
