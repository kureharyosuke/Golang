package main

import "fmt"

func main() {
	// 짧은 선언
	// 함수 안에서만 사용(전역X, Global No), 선언 후 할당 예외 발생
	// 주로 제한된 범위의 함수내에서 사용 할 경우 코드 가독성을 높일 수 있다. -> 추천

	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false

	// shortVar1 := 10  | 한번 할당하면, 할당 할수 없다.

	// shortVar3 := true // 예외발생

		fmt.Println("shortVar1 : ", shortVar1, "shortVar2 : ", shortVar2, "shortVar3" ,shortVar3 )

		var i int = 10
	// // Example
	if i < 15 {
		fmt.Println("i = 10")
	}

	 if i := 15; i < 20 {
		 // if i는 if에서만 쓴다고 생각할수 있다.
		 fmt.Println(("i = 15"))
	 }	
	}