// 변수
package main

import "fmt"

func main() {
	// 기본 초기화
	// 정수타입: 0, 실수(소수점): 0.0 문자열: "", Boolean: true, false
	// 변수명: 숫자첫글자X, 대문자구분 O, 문자, 숫자, _, 특수기호 사용가능
	// 변수 및 상수 : 함수 내외 사용 가능 (var, const)

	var a int
	var b string
	var c, d, e int
	var f, g, h int = 1, 2, 3
	var i float32 = 11.4
	var j string = "Hi! Golang"
	var k = 4.74 // 선언 동시 초기화
	var l = "Hi! Tokyo"
	var m = true

	a = 5
	b = "Hello Go"
	// c, d, e = 77, 100, 200
	c = 14
	d = 38
	e = 98

	fmt.Println("A : ", a)
	fmt.Println("B : ", b)
	fmt.Println("C, D, E  : ", c, d, e)
	fmt.Println("F, G, H : ", f, g, h)
	fmt.Println("I : ", i)
	fmt.Println("J : ", j)
	fmt.Println("K : ", k)
	fmt.Println("L : ", l)
	fmt.Println("M : ", m)
}