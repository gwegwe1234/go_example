package main

import "fmt"

func rangeSum(rg int, c chan int) {
	sum := 0

	for i := 0; i <= rg; i++ {
		sum += i
	}

	c <- sum
}

func main() {
	c := make(chan int)

	go rangeSum(1000, c)
	go rangeSum(700, c)
	go rangeSum(500, c)

	// 순서대로 데이터 수신(동기) : 채널에서 값 수신 완료 될 때까지 대기
	result1 := <-c
	result2 := <-c
	result3 := <-c

	fmt.Println("1 ", result1)
	fmt.Println("2, ", result2)
	fmt.Println("3, ", result3)
}
