// 열거형
// Iota -> enum class
// 상수 숫자 또는 규칙 나열할 때, 사용한다.

package main

import "fmt"

func main() {
	// 열거형
	// 상수를 사용하는 일정한 규칙에 따라 숫자를 계산 및 증가시키는 묶음
	
	const (
		Jan = 1
		Feb = 2
		Mar = 3
		Apr = 4
		May = 5
		Jun = 6

	)

	fmt.Println(Jan, Feb, Mar, Apr, May, Jun)
	
}