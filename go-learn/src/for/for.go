package main

import (
	"fmt"
)

func main() {
	// for = 반복문
	// go 에서 유일하게 제공되는 반복문
	// 다양한 방법 숙지

	fmt.Println("For Loop Start")
	// OUTPUT: For Loop Start

	// 예제 1
	for i := 0; i < 5; i++ {
		fmt.Println("ex1:", i)
	}
	// OUTPUT:

	// 자바스크립트
	// for (let = 0; i < 5; i++) {
	// 초기문 조건문 증강문
	// }

	// 에러발생 1 = 중가로 밑에 있는 경우
	/*
	 for i := 0; i < 5; i++      // <- 컴파일시,  i++에  ; 세미클론이 되기 때문에 에러난다.
	 {
		 fmt.Println("ex1 :", i)
	 }
	*/

	// 에러발생 2 = 중가로가  없는 경우
	/*
	 for i := 0; i < 5; i++      // <- 컴파일시,  i++에  ; 세미클론이 되기 때문에 에러난다.
		 fmt.Println("ex1 :", i)
	*/


	// 예제 2 ( 무한 루프)
	/*
	for {
		fmt.Println("ex2: Hello Golang")
		fmt.Println("ex2: Infinite Loop!")
	}
	*/

	// 예제 3 (Range 용법)
	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println("ex3: ", index, name)
	}
	/* OUTPUT
	ex3:  0 Seoul
	ex3:  1 Busan
	ex3:  2 Incheon
	*/

	for _, name := range loc {
		fmt.Println("ex4: ",  name )
	}
	/* OUT
	ex4:  Seoul
	ex4:  Busan
	ex4:  Incheon
	*/

	// _ = index 생략 무음

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum) 
	// OUTPUT: 5050


}