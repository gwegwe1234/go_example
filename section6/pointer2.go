package main

import "fmt"

func main() {
	i := 7
	p := &i

	fmt.Println(i, *p, &i, p)

	*p++

	fmt.Println(i, *p, &i, p)

	*p = 7777

	fmt.Println(i, *p, &i, p)

	i = 77

	fmt.Println(i, *p, &i, p)

	/*
		7 7 0xc0000180a8 0xc0000180a8
		8 8 0xc0000180a8 0xc0000180a8
		7777 7777 0xc0000180a8 0xc0000180a8
		77 77 0xc0000180a8 0xc0000180a8
	*/
}
