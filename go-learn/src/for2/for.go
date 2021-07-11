package main

import (
	"fmt"
)

func main() {
	// for = 반복문
	// go 에서 유일하게 제공되는 반복문
	// 다양한 방법 숙지

	fmt.Println("For Loop Start")
	


	// 예제 1

	var sum int = 0;
	// const sum int = 0;
	sum1 := 0;
	for i := 0; i <= 100; i++ {
		sum1 += 1 // 101
		sum += i	
	}
	fmt.Println("ex1 : ", sum1) //OUTPUT: ex1 :  101
	println("ex1-1:", sum) //OUTPUT: ex1-1: 5050


	
	// 예제 2

	sum2, i := 0, 0

	for i <= 100 {
		sum2 += i
		i++

		// j := i++ // 에러발생 Error: Go에서 후치연산은 반환값X
	}
	fmt.Println("ex2: ", sum2)


	
	// 예제 3
	sum3, i := 0, 0

	for { // while 형태랑 비슷!
		if i > 150 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("ex3: ", sum3)
	// OUTPUT: ex3:  11325




	
	// 예제 4
	for i, j := 0, 0; i <=10; i, j = i+1, j+10 {
		fmt.Println("ex4 :", i, j)
	}


	// 에러 발생 Error: 
	/*
		for i, j := 0, 0; i <=10; i, j = i++, j += 10 {	
			fmt.Println("ex4 :", i, j)
		}

		i++, j += 10 에러나는 부분, 후치연산은 안됨.
	*/


}