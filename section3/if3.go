package main

import "fmt"

func main()  {
	i := 100

	if i > 120 {
		fmt.Println("120 초과")
	} else if i > 100 && i < 120 {
		fmt.Println("100 초과 120 미만")
	} else {
		fmt.Println("100")
	}
}
