package main

import "fmt"

func main()  {
	var n1 uint8 = 125
	var n2 uint8 = 90

	fmt.Println(n1 + n2)
	fmt.Println(n1 - n2)
	fmt.Println(n1 * n2)
	fmt.Println(n1 / n2)
	fmt.Println(n1 % n2)
	fmt.Println(n1 << 2)
	fmt.Println(n1 >> 2)
	fmt.Println(^n1)

	/*
		215
		35
		242
		1
		35
		244
		31
		130
	 */

	 var n3 int = 12
	 var n4 float32 = 8.2
	 var n5 uint16 = 1024
	 var n6 uint32 = 120000

	 fmt.Println(float32(n3) + n4)
	 fmt.Println(n3 + int(n4))
	 fmt.Println(n5 + uint16(n6)) // 이상한 값 55488 나옴. 16으로 되면서 범위가 바뀌면서 숫자가 바뀜

	 /*
	    20.2
		20
		55488
	  */
}
