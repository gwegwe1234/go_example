package main

import "fmt"

func main() {
	// 맵
	// 맵 조회 및 순회

	// 예제 1

	map1 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google/com",
	}

	fmt.Println("ex1 : ", map1["google"])
	fmt.Println("ex1 : ", map1["daum"])
	fmt.Println("ex1 : ", map1["naver"])

	/*
		ex1 :  http://google/com
		ex1 :  http://daum.net
		ex1 :  http://naver.com
	*/

	// 예제 2

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	/*
		daum http://daum.net
		naver http://naver.com
		google http://google/com
	*/

	for _, v := range map1 {
		fmt.Println(v)
	}

	/*
		 	http://google/com
			http://daum.net
			http://naver.com
	*/

}
