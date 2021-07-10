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
	// "for 인덱스,요소값 := range 컬렉션"
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


	// 1. 1부터 100까지 더하는 for 루프문 
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum) 
	// OUTPUT: 5050


	// 2. 조건식 루프 for if
	n := 1
	for n < 100 {
		n *= 2
	}
	println(n)
	// OUTPUT: 128

	// 3. for range

	name := []string{"John", "OH", "Hayashi", "Ohhori"}

	for index, name := range name {
		println(index, name)
	}
	/*
	0 John
	1 OH
	2 Hayashi
	3 Ohhori
	*/

	// 4. break, continue, goto
	// break문은 for 루프 이외에 switch문이나 select문에서도 사용할 수 있다. 
	// 하지만, continue문은 for 루프와 연관되어 사용된다.
	
	var a int = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue // for 루프 시작으로
			// 만약 for 루프의 중간에서 나머지 문장들을 실행하지 않고 for 루프 시작부분으로 바로 가려면 continue문을 사용한다.
		}
		a++
		if a > 10 {
			break // 루프 빠져나옴
			// 경우에 따라 for 루프내에서 즉시 빠져나올 필요가 있는데, 이때 break 문을 사용한다. 
		}
	}
	if a == 11 {
		goto END // goto 사용예
		// 그리고 기타 임의의 문장으로 이동하기 위해 goto 문을 사용할 수 있다. goto문은 for 루프와 관련없이 사용될 수 있다.
	}
	println((a))

	END: 
	println("END")
	// OUTPUT: END


	/*
	break문은 보통 단독으로 사용되지만, 경우에 따라 "break 레이블"과 같이 사용하여 지정된 레이블로 이동할 수도 있다. 
	break의 "레이블"은 보통 현재의 for 루프를 바로 위에 적게 되는데,
	이러한 "break 레이블"은 현재의 루프를 빠져나와 지정된 레이블로 이동하고, break문의 직속 for 루프 전체의 다음 문장을 실행하게 한다.
	*/
	
	i := 0
	
	L1:
    for {
     
        if i == 0 {
            break L1
        }
    }
 
    println("OK")
	// 언뜻 보기에 무한루프를 돌 것 같지만, 실제로는 OK를 출력하고 프로그램을 정상 종료한다
	// 이는 "break L1" 문이 for 루프를 빠져나와 L1 레이블로 이동한 후, 
	// break가 있는 현재 for 루프를 건너뛰고 다음 문장인 println() 으로 이동하기 때문이다.
}