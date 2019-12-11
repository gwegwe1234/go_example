package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) Caculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	// 구조체
	// GO의 필드들의 집합체 또는 컨테이너
	// 필드는 갖지만 메소드는 갖지 않는다.
	// 객체지향 방식을 지원 -> 리시버를 통해 메소드랑 연동
	// 상속, 객체, 클래스 개념 없음
	// 구조체 -> 구조체 선언 -> 구조체 인스턴스(리시버)

	kim := Account{number: "245-901", balance: 10000000, interest: 0.5}
	lee := Account{number: "245-123", balance: 40101111100}
	park := Account{number: "245-124", interest: 0.4}
	choi := Account{"245-123", 1500000, 0.3}

	fmt.Println(kim)
	fmt.Println(lee)
	fmt.Println(park)
	fmt.Println(choi)

	// 예제 2
	fmt.Println(int(kim.Caculate()))
	fmt.Println(int(lee.Caculate()))
	fmt.Println(int(park.Caculate()))
	fmt.Println(int(choi.Caculate()))
}
