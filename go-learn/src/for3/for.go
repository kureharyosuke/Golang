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
	Loop1:
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if i == 2 && j == 4 {
					break Loop1
				}
				fmt.Println("ex1 : ", i, j) 
			}
		} 
//OUTPUT
// ex1 :  0 0
// ex1 :  0 1
// ex1 :  0 2
// ex1 :  0 3
// ex1 :  0 4
// ex1 :  1 0
// ex1 :  1 1
// ex1 :  1 2
// ex1 :  1 3
// ex1 :  1 4
// ex1 :  2 0
// ex1 :  2 1
// ex1 :  2 2
// ex1 :  2 3

	for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if i == 2 && j == 4 {
					break
				}
				fmt.Println("ex1 : ", i, j) 
			}
		} 
	//OUTPUT
	// ex1 :  0 0
	// ex1 :  0 1
	// ex1 :  0 2
	// ex1 :  0 3
	// ex1 :  0 4
	// ex1 :  1 0
	// ex1 :  1 1
	// ex1 :  1 2
	// ex1 :  1 3
	// ex1 :  1 4
	// ex1 :  2 0
	// ex1 :  2 1
	// ex1 :  2 2
	// ex1 :  2 3
	// ex1 :  3 0
	// ex1 :  3 1
	// ex1 :  3 2
	// ex1 :  3 3
	// ex1 :  3 4
	// ex1 :  4 0
	// ex1 :  4 1
	// ex1 :  4 2
	// ex1 :  4 3
	// ex1 :  4 4


	
	// 예제 2
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue // for 루프 시작부분으로 바로 가려면 continue문을 사용한다.  
			// 여기서 2의 배수의 조건은 fmt.Println으로 못가고 for문으로 다시 돌아간다.
		}
		fmt.Println("ex2 : ", i)
	}


	
	Loop2:
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if i == 1 && j == 2 {
					continue Loop2
				}
				fmt.Println("ex3 : ", i, j )
			}
		}



}