// 열거형
// Iota -> enum class
// 상수 숫자 또는 규칙 나열할 때, 사용한다.

package main

import "fmt"

func main() {
	// 열거형
	// 상수를 사용하는 일정한 규칙에 따라 숫자를 계산 및 증가시키는 묶음
	
	const (
		A = iota
		B 
		C
		D
	)

	fmt.Println(A, B, C, D) //0 1 2 3

	const (
		F = iota * 10
		G 
		H
		J
	)

	fmt.Println(F, G, H, J) //0 10 20 30

	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
		May
		Jun

	)

	fmt.Println(Jan, Feb, Mar, Apr, May, Jun) // 1 2 3 4 5 6
	
}