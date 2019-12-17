package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 뮤텍스 : 상호 배제 -> Thread(고루틴)들이 서로 running time에 서로 영향을 주지 않게 즉, 단독으로 실행되게 하는 기술
	// 뮤텍스 : 여러 고루틴에서 작업하는 공유 데이터 보호

	// 동기화 사용하지 않는 경우 예제
	// 쓰기 읽기 동작 순서가 일정하지 않아 잘못된 오류를 반환 할 가능성 증가

	// CPU 전체사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0

	go func() {
		for i := 1; i <= 10; i++ {
			data += 1
			fmt.Println("Write : ", data)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Read1 : ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Read2 : ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(10 * time.Second)

	/*
		Write :  1
		Read2 :  1
		Read1 :  1
		Write :  2
		Write :  3
		Write :  4
		Write :  5
		Read2 :  5
		Read1 :  5
		Write :  6
		Write :  7
		Write :  8
		Write :  9
		Write :  10
		Read2 :  10
		Read1 :  10
		Read2 :  10
	*/
	// 이상하게 1쓰고 1읽고, 2쓰고 2 읽고 이런식으로 동기화가 제대로 안됐다.

}
