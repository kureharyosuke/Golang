package main

import (
	"fmt"
	"math/rand" // 숫자의 난수
	"time"      // 시간에 대한 타임 패키지
)

func main() {


	// rand.Seed() == 숫자의 난수
	// time.Now() == 현재의 시간
	// UnixNano() == 유닉스 나노의 기준으로 가져오겠다.
	// rand.Intn == rand의 숫자

	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("i -> ", i, " 50 이상 100 미만")
	case i >= 25 && i < 50:
		fmt.Println("i -> ", i, " 25 이상 100 미만")
	default: 
		fmt.Println("i -> ", i, " 기본 값")
	}
	// OUTPUT: i ->  68  50 이상 100 미만
	// OUTPUT2: i ->  7  기본 값
}