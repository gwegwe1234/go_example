package main

// Alias 지정

import (
	"fmt"
	checkUp "github.com/gwegwe1234/go_example/section4/lib"
	checkUp2 "github.com/gwegwe1234/go_example/section4/lib2"
	// _ "github.com/gwegwe1234/go_example/section4/lib2" -> 빈 식별자. _ 를 쓰면 ide에서도 안지움
)

func main()  {
	// 패키지 접근제어
	// 별칭 사용
	// 빈 식별자 사용

	fmt.Println("10보다 큰 수? : ", checkUp.CheckNum(20))
	fmt.Println("100보다 큰 수? : ", checkUp2.CheckNum1(200))
}
