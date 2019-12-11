package main

import "fmt"

type car struct {
	color string
	name  string
}

func main() {
	// 예제 1
	c1 := car{"Red", "220d"}
	c2 := new(car)
	c2.color, c2.name = "Yellow", "sonata"
	c3 := &car{}
	c4 := &car{"black", "520d"}

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)

	/*
		{Red 220d}
		&{Yellow sonata}
		&{ }
		&{black 520d}
	*/
}
