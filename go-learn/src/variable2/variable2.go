package main

import "fmt"

func main() {
	//여러개 선언	
	var (
		name string = "machine"
		height int32
		weight float32
		isRunning bool
		// 가독성과, 그룹을 묶어 하나의 상태를 의미
	)

	height = 250
	weight = 350.56
	isRunning = true

	fmt.Println("name : ", name, "height : ", height, "weight : ", weight, "isRunning : ", isRunning)
}