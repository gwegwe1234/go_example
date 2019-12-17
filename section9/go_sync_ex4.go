package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	// 원자성 사용 -> 기능적으로 분할 불가능한 완전 보증된 일련의 조작, 모두 성공하거나 모두 실패하거나
	// 모든 조작이 완료 될 때까지 다른 프로세스 개입 불가
	// sync/atomic 에서 원자적 연산자 제공
	// https://golang.org/pkg/sync/atomic
	// 주로 공용 변수에 관한 계산 사용

	// 원자성 사용 할 경우 예제
	runtime.GOMAXPROCS(runtime.NumCPU())

	var cnt int64 = 0
	waitGroup := new(sync.WaitGroup)

	for i := 0; i < 5000; i++ {
		waitGroup.Add(1)
		go func(n int) {
			//cnt += 1
			atomic.AddInt64(&cnt, 1)
			waitGroup.Done()
		}(i)
	}

	for i := 0; i < 2000; i++ {
		waitGroup.Add(1)
		go func(n int) {
			//cnt -= 1
			atomic.AddInt64(&cnt, -1)
			waitGroup.Done()
		}(i)
	}

	// Add == Done
	waitGroup.Wait()

	finalCnt := atomic.LoadInt64(&cnt)
	fmt.Println("Wait Group End! >>>>>> ", cnt)
	fmt.Println("Wait Group End! >>>>>> ", finalCnt)
	// Wait Group End! >>>>>>  3000
	// Wait Group End! >>>>>>  3000
}
