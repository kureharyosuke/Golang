package main

import "fmt"

func main() {
	// const 상수
	// const 사용 초기화, 한번 선언 후 값 변경금지, 고정된 값 관리용
	// const 값을 지정해야한다.
	
	const a string = "test1"
	const b = "Test2"
	const c int32 = 10 * 10
	// const d = getHeight() 상수에서는 함수를 사용할수 없다.

	const e = 35.6 // floot
	const f = false // bool
	
	/*
		에러 발생
		const g string
		g = "Test3"
	*/

	fmt.Println("a :", a, "b :", b, "c :", c, "e :", e, "f :", f)
	//a : test1 b : Test2 c : 100 e : 35.6 f : false
}