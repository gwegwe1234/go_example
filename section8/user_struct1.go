package main

import "fmt"

type Car struct {
	name  string
	color string
	price int64
	tax   int64
}

// 일반 메소드 -> 이렇게는 잘 안씀
func Price(c Car) int64 {
	return c.price + c.tax
}

// 구조체 <-> 메소드 바인딩
func (c Car) Price() int64 {
	return c.price + c.tax
}

func main() {
	// Go -> 객체 지향 타입을 구조체로 정의한다. (클래스, 상속 개념 없음)
	// 객체지향 -> 클래스(속성 : 멤버변수, 기능(상태: 메소드)) : 코드의 재사용성, 코드의 관리가 용이, 신뢰성이 높은 프로그래밍
	// 객체 지향 언어일까?
	// Go는 전형적인 객체지향의 특징을 가지고 있지 않지만, 인터페이스 등을 통한 다형성 지원, 구조체를 통한 클래스 형태의 코딩
	// 객체지향의 기본 개념은 Go에서 갖고있음
	// 상태, 메소드 분리해서 정의 (결합성 없음)
	// 사용자 정의타입 : 구조체, 인터페이스, 기본타입, 함수

	// 예제 1
	bmw := Car{name: "520d", price: 50000000, color: "white", tax: 5000000}
	benz := Car{name: "220d", price: 60000000, color: "white", tax: 6000000}

	fmt.Println("ex1 : ", bmw, &bmw)
	fmt.Println("ex1 : ", benz, &benz)

	/*
		ex1 :  {520d white 50000000 5000000} &{520d white 50000000 5000000}
		ex1 :  {220d white 60000000 6000000} &{220d white 60000000 6000000}
	*/

	fmt.Println("ex2 : ", Price(benz))
	fmt.Println("ex2 : ", Price(bmw))

	fmt.Println("ex3 : ", bmw.Price(), benz.Price())

	fmt.Println("ex4 : ", &bmw == &benz)
}
