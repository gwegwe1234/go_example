package main

import (
	"fmt"
	"sync"
)

func main() {
	// 대기 그룹
	// 실행 흐름 변경 (고루틴이 종료 될 때 까지 대기 가능)
	// WaitGroup : Add(고루틴 추가), Done(작업 종료 알림), Wait(고루틴 종료시 까지 대기)
	// Add로 추가 된 고루틴 개수와 Done 으로 종료되는 알림 횟수가 동일해야 정상적으로 작

	waitGroup := new(sync.WaitGroup)
	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func(n int) {
			fmt.Println("Wait Group: ", n)
			waitGroup.Done()
		}(i)
	}

	// Add == Done
	waitGroup.Wait()
	fmt.Println("Wait Group End!") // 훨씬 깔
}
