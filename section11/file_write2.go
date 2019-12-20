package main

import (
	"encoding/csv"
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
	// CSV 파일 쓰기 예제
	// 패키지 저장소를 통해서 Excel 등 다양한 파일 형식 쓰기, 읽기 가능
	// 패키지 github 주소 : https://github.com/tealeg/xlsx
	// bufio : 파일이 용량이 클 경우 버퍼 사용 권장

	file, err := os.Create("./src/github.com/gwegwe1234/go_example/section11/test_write.csv")
	errCheck1(err)

	defer file.Close()

	wr := csv.NewWriter(file)
	// wr := csv.NewWriter(bufio.NewWriter(file)) 권장

	wr.Write([]string{"kim", "4.8"})
	wr.Write([]string{"park", "4.5"})
	wr.Write([]string{"lee", "4.8"})
	wr.Write([]string{"kkk", "4.8"})
	wr.Write([]string{"kim", "4.8"})
	wr.Write([]string{"kim", "4.8"})

	wr.Flush()

	fi, err := file.Stat()
	errCheck1(err)

	fmt.Printf("CSV 쓰기 작업 후 파일 크기(%d bytes)\n", fi.Size())
	fmt.Println("CSV 파일명 : ", fi.Name())
	fmt.Println("운영체제 파일 권한 : ", fi.Mode())
}
