package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func CaculateD(a Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CaculateP(a *Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func main() {
	kim := Account{"254-021", 1000, 1}
	lee := Account{"254-021", 2000, 2}

	fmt.Println(kim, lee)

	CaculateD(kim)
	CaculateP(&lee)

	fmt.Println(kim.balance, lee.balance)
}
