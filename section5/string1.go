package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {
	// 문자열
	// 큰 따옴표 "", 백스쿼트 ``
	// 문자 char '' 타입 존재 하지 않음 -> rune(int32)로 문자 코드 값으로 표현
	// 문자 : '' 로 작성
	// 자주 사용하는 escape : \\, \', \", \a (콘솔 벨), \b (백스페이스), \f (쪽바굼), \n (줄바꿈)

	// 예제 1
	var str1 string = "~\\"
	str2 := `~\` // 이스케이프 사용 안하고 백스쿼트를 사용하면 좀더 깔끔하다

	fmt.Println(str1)
	fmt.Println(str2)

	// 예제 2
	var str3 string = "Hello World!"
	var str4 string = "안녕하세요."

	fmt.Println(str3)
	fmt.Println(str4)

	// 예제 3
	// 길이 (바이트수)
	fmt.Println(len(str3))
	fmt.Println(len(str4))

	// 예제 4
	// 실제 길이 (utf8에서 제공)
	fmt.Println(utf8.RuneCountInString(str3))
	fmt.Println(utf8.RuneCountInString(str4))
	fmt.Println(len([]rune(str4))) // Len으로 실제 길이 구하는 법 but 보통 utf8을 사용한다
}
