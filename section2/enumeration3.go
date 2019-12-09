package main

import "fmt"

func main() {
	const (
		_ = iota
		A
		_
		C
	)

	const (
		_ = iota + 0.75*2
		DEFAULT
		SILVER
		GOLD
		PLATINUM
	)

	fmt.Println(A, C)
	// 1 3

	fmt.Println(DEFAULT, SILVER, GOLD, PLATINUM)
	// 2.5 3.5 4.5 5.5
}
