package main

import "fmt"

func init() {
	// main 보다 먼저 실행된다.
	// init : 패키지 로드 시에 가장 먼저 호출
	// 가장 먼저 초기화 되는 작업 작성 시 유용하다.
	fmt.Println("Init Method Start")
}

func main() {
	fmt.Println("Main Method Start")
}
