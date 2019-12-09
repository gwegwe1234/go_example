package main

import (
	"fmt"
	"github.com/gwegwe1234/go_example/section4/lib"
)

var num int32

func init() {
	num = 30
	fmt.Println("Main Init Start")
}

func main() {
	fmt.Println("10보다 큰수 ? : ", lib.CheckNum(num))
}
