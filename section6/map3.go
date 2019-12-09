package main

import "fmt"

func main() {

	// 예제 3
	// 맵 값 변경 및 삭제

	map1 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google/com",
		"home1":  "http://test.com",
	}

	fmt.Println("ex1 : ", map1)
	// ex1 :  map[daum:http://daum.net google:http://google/com home1:http://test.com naver:http://naver.com]

	map1["home2"] = "http://test2.com" // 추가
	fmt.Println("ex1 : ", map1)
	// ex1 :  map[daum:http://daum.net google:http://google/com home1:http://test.com home2:http://test2.com naver:http://naver.com]

	map1["home2"] = "http://test3.com" // 변경
	fmt.Println("ex1 : ", map1)
	// ex1 :  map[daum:http://daum.net google:http://google/com home1:http://test.com home2:http://test3.com naver:http://naver.com]

	// 예제 4

	delete(map1, "home2")
	fmt.Println("ex2 : ", map1)
	// ex2 :  map[daum:http://daum.net google:http://google/com home1:http://test.com naver:http://naver.com]

	delete(map1, "home1")
	fmt.Println("ex2 : ", map1)
	// ex2 :  map[daum:http://daum.net google:http://google/com naver:http://naver.com]
}
