package main

import "fmt"

func main()  {
	// 여러개 설정
	var (
		name string = "machine"
		height int32
		weight float32
		isRunning bool
	)

	height = 250
	weight = 350.56
	isRunning = true

	fmt.Println("name : ", name)
	fmt.Println("height : ", height)
	fmt.Println("weight : ", weight)
	fmt.Println("isRunning : ", isRunning)
}
