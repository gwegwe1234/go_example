package main

import "fmt"

func main() {
	// 구조체 익명 선언

	car1 := struct{ name, color string }{"520d", "red"}

	fmt.Println("Ex1 : ", car1)
	fmt.Printf("ex1 : %#v", car1)
	fmt.Println()
	cars := []struct{ name, color string }{{"530d", "red"}, {"540d", "blue"}, {"550d", "green"}}

	for _, c := range cars {
		fmt.Println(c)
	}
}
