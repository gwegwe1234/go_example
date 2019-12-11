package main

import "fmt"

type shoppingBasket struct{ cnt, price int }

func (b shoppingBasket) purchase() int {
	return b.cnt * b.price
}

// 원본이 수정 x (참조 형식)
func (b shoppingBasket) rePurchaseD(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

// 원본이 수정 (참조 형식)
func (b *shoppingBasket) rePurchaseP(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func main() {
	// 리시버 (구조체 메소드) 전달(값, 참조) 형식
	// 함수는 기본적으로 값 호출 -> 변수으 값이 복사 후 내부 전달(원본 수정 x) -> 맵, 슬라이스는 참조 전달
	// 구조체도 참조 전달..!
	// 근데 따로 역참조 하고 이럴필요가 없다. * & 이런거 필요없음. (..)

	// 예제 1

	bs1 := shoppingBasket{3, 5000}
	fmt.Println("ex1(totPrice) : ", bs1.purchase())

	bs1.rePurchaseP(7, 5000)
	fmt.Println("ex2(totPrice) : ", bs1.purchase())

	bs1.rePurchaseD(5, 5000)
	fmt.Println("ex3(totPrice) : ", bs1.purchase())
}
