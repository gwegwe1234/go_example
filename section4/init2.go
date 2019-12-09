package main

import "fmt"

func init() {
	// init 만 있으면 당연히 안됨.
	// main 이 없으면 library라고 보면된다. 다른 메인에서 갖다 써야됨
	fmt.Println("Init Method Start")
}
