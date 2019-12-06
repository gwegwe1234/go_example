package main

import (
	"fmt"
	"os"
)

// 선언방법 1
/*
import "fmt"
import "os"
*/

// 선언방법 2
/*
import (
	"fmt"
	"os"
)
 */

func main()  {
	// 패키지 : 코드 구조화 (모듈) 및 재사용
	// 응집도 up, 결합도 down
	// Go : 패키지 단위의 독립적이고 작은 단위로 개발 -> 작은 패키지를 결합해서 프로그램을 작성하는걸 권고
	// 패키지 이름 = 디렉토리 이름
	// 같은 패키지 내 -> 소스파일들은 디렉토리명을 패키지 명으로 사용한다.
	// 네이밍 규칙 : 소문자 private, 대문자 public
	// fmt 같은 패키지는 go env 내의 GOROOT 경로에 있다.
	// Go : main 패키지는 특별하게 인식 -> 컴파일러 공유 라이브러리 x, 프로그램의 시작점 start point로 인식한다.

	var name string

	fmt.Println("이름은 ? : ") // Println 이 대문자, 즉 퍼블릭이다.

	fmt.Scanf("%s", &name)

	fmt.Fprintf(os.Stdout, "Hi! %s\n", name)

	/*
		이름은 ? :
		박광태
		Hi! 박광태
	 */
}
