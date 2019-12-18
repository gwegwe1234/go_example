package main

import (
	"fmt"
	"os"
)

func fileOpne(filename string) {
	// defer 함수 (Panic 호출되면 실행)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("File Open Error : ", r)
		}
	}()

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("file! open!!!")
	}

	defer f.Close()
}

func main() {
	fileOpne("undefined.txt")
	fmt.Println("End Main")
}
