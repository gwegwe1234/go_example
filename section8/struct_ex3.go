package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) CaculateD(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a *Account) CaculateP(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func main() {
	// 정리 : 구조체 인스턴스 값 변경시 -> 포인터 전달, 보통의 경우 -> 값전달
	kim := Account{"254-021", 1000, 1}
	lee := Account{"254-021", 2000, 2}

	lee.CaculateD(400)
	kim.CaculateP(10100)
	fmt.Println(kim, lee)

}
