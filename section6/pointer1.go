package main

import "fmt"

func main() {
	// 포인터
	// Go : 포인터 지원(c)
	// 변수의 지역성, 연속된 메모리 참조성 , 힙영역, 스택..
	// 파이썬, 자바는 컴파일러, 인터프리터 내부에서 쓰긴한다.
	// 포인터 지원 X (파이썬, C#, JAVA 등)
	// 주소의 값은 직접 변경 불가능 (잘못된 코딩으로 인한 버그 방지)
	// *
	// nil로 초기화 (nil != 0)

	// 예제 1
	var a *int            // 방법 1
	var b *int = new(int) // 정석

	fmt.Println(a)
	fmt.Println(b)

	/*
		<nil>
		0xc0000aa000
	*/

	i := 7

	fmt.Println("ex1 : ", i, &i)
	// ex1 :  7 0xc0000ac000

	a = &i // 주소값 전달
	b = &i

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(*a) // 역참조

	fmt.Println()

	fmt.Println(b)
	fmt.Println(&b)
	fmt.Println(*b)

	/*
		 	0xc0000ac010
			0xc00000e020 // a 자체의 주소
			7

			0xc0000ac010
			0xc00000e028 // b 자체의 주소
			7
	*/

	var c = &i
	d := &i

	*d = 10

	fmt.Println(c)
	fmt.Println(&c)
	fmt.Println(*c) // 역참조

	fmt.Println()

	fmt.Println(d)
	fmt.Println(&d)
	fmt.Println(*d)
	/*
		0xc0000aa010
		0xc0000b6000
		10 // 싹다 다바뀜

		0xc0000aa010
		0xc0000b6008
		10
	*/
}
