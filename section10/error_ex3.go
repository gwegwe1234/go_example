package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

// 예외 처리 구조체
type PowError struct {
	time    time.Time // 발생 시간
	value   float64   // 파라미터
	message string    // 에러 메시지
}

func (e PowError) Error() string {
	return fmt.Sprintf("[%v]Error - Input Value(value : %g) - %s", e.time, e.value, e.message)
}

func Power(f, i float64) (float64, error) {
	if f == 0 {
		return 0, PowError{time: time.Now(), value: f, message: "0은 사용할 수 없습니다"}
	}

	if math.IsNaN(f) {
		return 0, PowError{time: time.Now(), value: f, message: "문자는 사용할 수 없습니다"}
	}

	if math.IsNaN(i) {
		return 0, PowError{time: time.Now(), value: f, message: "문자는 사용할 수 없습니다"}
	}

	return math.Pow(f, i), nil
}

func main() {
	//error 타입이 아닌경우 에러 처리 방법
	// Error 메소드를 구현해서 사용자 정의 에러 처리 예제 심화
	// 구조체를 사용해서 세부적인 정보 출력

	v, err := Power(10, 3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)

	t, err := Power(0, 3)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err.(PowError))
	}

	fmt.Println(t)

}
