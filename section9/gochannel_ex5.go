package main

import (
	"fmt"
	"time"
)

func main() {
	// 채널

	// 예제 1 발신 스위치문 처리
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			num := <-ch1 // 값 수신
			fmt.Println("ch1 : ", num)
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Hi Golang!"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case ch1 <- 777:
			case str := <-ch2:
				fmt.Println("ch2 : ", str)
			}
		}
	}()
	time.Sleep(7 * time.Second)
}
