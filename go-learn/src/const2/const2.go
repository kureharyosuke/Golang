package main

import "fmt"

func main() {
	const a, b int = 10, 99
	const c, d, e  = true, 0.83, "Test"
	
	const (
		x, y int16 = 50, 90
		i, k       = "Data", 7776
	)

	fmt.Println(a,b,c,d,e) //OUT: 10 99 true 0.83 Test
	fmt.Println(x,y,i,k) //OUT: 50 90 Data 7776
}