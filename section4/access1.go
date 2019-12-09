package main

import (
	"fmt"
	"github.com/gwegwe1234/go_example/section4/lib2"
)

func main() {
	// 패키지 접근제어
	// 변수, 상수, 함수, 메소드, 구조체 들 식별자
	// 대문자 : 패키지 외부에서 접근 가능
	// 소문자 : 패키지 외부 점근 불가

	fmt.Println("100보다 큰수 ? ", lib2.CheckNum1(101))
	fmt.Println("1000보다 큰수 ? ", lib2.CheckNum2(999))
}
