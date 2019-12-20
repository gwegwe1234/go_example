package main

import (
	"fmt"
	"os"
)

func errCheck1(e error) {
	if e != nil {
		panic(e)
	}
}

func errCheck2(e error) {
	if e != nil {
		fmt.Println("Error!")
	}
}

func main() {
	// 파일 쓰기
	// Create : 새 파일 작성 및 파일 열기
	// Close : 닫기
	// Write, WriteString, WriteAt : 파일 쓰기
	// 각 운영체제 권한 주의()
	// 예외 처리 중요

	// 파일 쓰기 예제

	file, err := os.Create("./src/github.com/gwegwe1234/go_example/section11/test_write.txt")
	errCheck1(err)

	defer file.Close()

	s1 := []byte{115, 111, 109, 101, 115}

	n1, err := file.Write([]byte(s1)) // 문자열 -> byte 슬라이스 형으로 변환 후 쓰기
	errCheck2(err)

	fmt.Printf("쓰기 작업(1) 완료 (%d bytes) \n", n1)

	file.Sync() // Write Commit (stable)

	// string 쓰기 예제

	s2 := "Hello Golang! \n File Write Test -1 \n"
	n2, err := file.WriteString(s2)
	errCheck2(err)

	fmt.Printf("쓰기 작업(2) 완료 (%d bytes)\n", n2)

	// 예제 3

	s3 := "Test Write At! -2 \n"
	n3, err := file.WriteAt([]byte(s3), 50)
	errCheck2(err)

	fmt.Printf("쓰기 작업(3) 완료 (%d bytes)\n", n3)

	file.Sync()

	// 예제 4

	n4, err := file.WriteString("Hello Golang 2!!!\n")
	errCheck2(err)

	fmt.Printf("쓰기 작업(4) 완료 (%d bytes)\n", n4)
}
