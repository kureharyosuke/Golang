// 열거형
// Iota -> enum class
// 상수 숫자 또는 규칙 나열할 때, 사용한다.

package main

import "fmt"

func iotafuc() {
	// 열거형
	// 상수를 사용하는 일정한 규칙에 따라 숫자를 계산 및 증가시키는 묶음
	
	const (
		_ = iota
		A
		_
		C
		E
	)
	
	fmt.Println(A, C, E) // 1 3 4

	const (
		_ = iota + 0.75 * 2
		DEFAULT // = 1 + 0.75 * 2
		SILVER // = 2 + 0.75 * 2
		GOLD // = 3 + 0.75 * 2
		PLATINUM // = 4 + 0.75 * 2
	)

	fmt.Println("D : ", DEFAULT)
	fmt.Println("S : ", SILVER)
	fmt.Println("G : ", GOLD)
	fmt.Println("P : ", PLATINUM)

/*
	D :  2.5
	S :  3.5
	G :  4.5
	P :  5.5
 */
}