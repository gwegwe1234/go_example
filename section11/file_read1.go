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
	// 파일 읽기
	// Open : 기존 파일 열기
	// Close : 리소스 닫기
	// Read, ReadAt
	// 각 운영체제 권한 주의
	// 예외처리 중요
	// 탐색 Seek 중요
	// 파일 읽기 예제
	// 파일 열기
	file, err := os.Open("./src/github.com/gwegwe1234/go_example/section11/readsample.txt")

	fileInnfo, err := file.Stat()
	errCheck1(err)

	fd1 := make([]byte, fileInnfo.Size()) // 슬라이스에 읽은 내용 담는다.
	ct1, err := file.Read(fd1)

	fmt.Println(fileInnfo)
	fmt.Println(fileInnfo.Name())
	fmt.Println(fileInnfo.Size())
	fmt.Println(fileInnfo.ModTime())
	fmt.Println(ct1)
	fmt.Println(fd1)
	fmt.Println(string(fd1))

	fmt.Println("==================================================================")

	// 읽기 예제 2 Seek (Offset)
	o1, err := file.Seek(5, 0) // 0: 처음 위치, 1 : 현재 오프셋, 2 : 끝
	errCheck1(err)

	fd2 := make([]byte, 5)
	ct2, err := file.Read(fd2)
	fmt.Println(ct2)
	fmt.Println(o1)

	fmt.Println(string(fd2))

	//
	o2, err := file.Seek(0, 0)
	errCheck1(err)

	fd3 := make([]byte, 50)
	ct3, err := file.ReadAt(fd3, 8)

	fmt.Println(ct3)
	fmt.Println(o2)

	fmt.Println(string(fd3))

	file.Close()

}
