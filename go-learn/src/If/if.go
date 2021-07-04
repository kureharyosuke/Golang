// IF 문

package main

import (
	"fmt"
)

func main() {
	// 조건문 (제어문)
	// IF 문 : 반드시 Boolean 검사 !== 1, 0 으로 사용할수 없다.(사용불가: 자동 형 반환 불가)
	// 소괄호 사용 X

	 var a int = 20
	 b:= 20

	 // 예제1
	 if a >= 15 {
		 fmt.Println("a는 15이상이다.")
	 }

	 //OUTPUT: a는 15이상이다.

	 // 예제2
	 if b >= 25 {
		 fmt.Println("b는 25이상이다.")
	 }

	 //OUTPUT: X b <= 25


	 // #Error1 에러발생 유형
	 
	 /*
	 if b >= 25;
	 {

	 }
	 */

	 // #Error2 에러발생 유형
	 
	 /*
	 if b >= 25
	 		fmt.Println("25이상")
		*/


	 // #Error3 에러발생 유형  0, 1 로 true false를 사용하면 안된다. 
	 
	 /* 불리언값으로 해라
		if c:= 1; c { // non-boolean
			fmt.Println("True")
		}
	 */


		if c:= true; c { // 사용가능
			fmt.Println("True")
		}
	 
		// 예제3
		if c := 40; c >= 35 {
			fmt.Println(("c는 35이상이다."))
		}

		// c += 20 
		// // OUTPUT : c는 undefind , 위의 c는 if문을 사용하고 소멸된 c 이기 때문에!

}
