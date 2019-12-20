package main

import (
	"bufio"
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
	file, err := os.Open("./src/github.com/gwegwe1234/go_example/section11/test_write.csv")

	errCheck1(err)
	defer file.Close()

	// CSV Reader 생성
	rr := csv.NewReader(bufio.NewReader(file))

	row, err := rr.Read()
	errCheck1(err)

	fmt.Println(row)

	// 한번에 갖오기

	rows, err := rr.ReadAll()
	errCheck1(err)
	fmt.Println(rows)
}
