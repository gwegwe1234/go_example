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
	// 다양한 선언방법
	// &struct, struct : &struct 포인터를 받아오고 역참조를 또 하기 때문에 속도는 조금 느리다.
	// 인터페이스 메소드를 선언만 해둔 후 -> 오버라이딩 해서 메소드에 포인터 리시버를 사용할 경우 반드시  &struct로 넘겨야하마

	// 선언방법 1
	var kim *Account = new(Account) // *구조체 =  new 사용
	kim.number = "245-901"
	kim.balance = 10000000
	kim.interest = 0.5

	// 선언방법 2
	hong := &Account{"245-902", 1000000, 0.4}

	lee := new(Account)
	lee.number = "245-903"
	lee.balance = 1000000
	lee.interest = 4

	fmt.Println(hong)
	fmt.Println(lee)
	fmt.Println(kim)

	// 예제 2

	fmt.Println("ex2 : ", int(kim.Caculate()))
	fmt.Println("ex2 : ", int(hong.Caculate()))
	fmt.Println("ex2 : ", int(lee.Caculate()))
}
