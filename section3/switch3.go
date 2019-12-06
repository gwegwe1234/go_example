package main

import "fmt"

func main()  {
	a := 30 / 15
	switch a {
	case 2, 4, 6: // i 가 2, 4, 6 일 경우
		fmt.Println("a -> ", a , " 는 짝수")
	case 1, 3, 5:
		fmt.Println("a -> ", a , " 는 홀수")
	}

	switch e := "go"; e {
	case "java":
		fmt.Println("Java!")
		fallthrough
	case "go":
		fmt.Println("go!")
		fallthrough
	case "python":
		fmt.Println("python!")
		//fallthrough
	case "c":
		fmt.Println("c!")
		fallthrough
	case "c#":
		fmt.Println("c#!")
	}
	// 결과 :
	// go!
	// python!
}