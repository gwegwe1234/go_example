package main

import (
	"errors"
	"fmt"
	"log"
)

func notZero(n int) (string, error) {
	if n != 0 {
		s := fmt.Sprint("hello golang; ", n)
		return s, nil
	}

	return "", errors.New("0을 입력했습니다. 에러 발생!")
}

func main() {
	// Errorf를 이용한 에러 처리

	a, err := notZero(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ex1 : ", a)

	b, err := notZero(0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ex1 : ", b)

	fmt.Println("End Error Handling!")
}
