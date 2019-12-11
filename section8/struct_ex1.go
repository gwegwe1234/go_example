package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func NewAccount(number string, balance float64, interest float64) *Account {
	return &Account{number, balance, interest} // 구조체 생성한 뒤 포인터 리턴
}

func main() {
	// 구조체 생성자 패턴 예제

	// 예제 1
	kim := Account{"245-000", 100000, 4}

	var lee *Account = new(Account)
	lee.number = "123-456"
	lee.balance = 123123123
	lee.interest = 3

	park := NewAccount("245-934", 10000, 10)
	park.interest = 4
	fmt.Println(kim, lee, park)
}
