package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 에러 처리
	// 에러 처리 : 소프트웨어의 품질 향상에 가장 중요한 것 => 유형코드 및 에러 정보 등을 정보를 남기는 것
	// Go에서는 기본적으로 error 패키지에서 error 인터페이스를 제공
	// type error interface { Error() string }
	// 즉, Error 메서드를 구현하면 사용자 저으이 에러 처리 제작 가능
	// 기본적으로 메소드 마다 리턴 타입 2개 (리턴값, 에러)
	// 주로 Errorf(에러 타입 리턴) 메소드, Fatal(프로그램 종료) 메소드를 통해서 에러 출력

	// 기본적인 메소드 에러 처리 예제
	f, error := os.Open("unnamedfile") // error 가 nil이면 에러가 안난거

	if error != nil {
		log.Fatal(error.Error()) // 방법 1
		//log.Fatal(error) // 방법 2
	}

	fmt.Println("==============")
	fmt.Println("ex1 : ", f.Name())

}
