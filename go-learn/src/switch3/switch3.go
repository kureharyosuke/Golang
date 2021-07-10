package main

import (
	"fmt"
)

func main() {

	// 예제 1
	a := 30 / 15
	switch a {
	case 2, 4, 6: // i가 2, 4, 6인 경우
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5:
		fmt.Println("a -> ", a, "는 홀수")
	default:
		fmt.Println("숫자가 아닙니다.")
	}
	// OUTPUT: a ->  2 는 짝수


	// 예제 2
	switch e := "go"; e {
	case "java":
		fmt.Println("java!")
		fallthrough
	case "go":
		fmt.Println("go!")
		fallthrough
	case "python":
		fmt.Println("python!")
		fallthrough
	case "ruby":
		fmt.Println("ruby!")
		fallthrough
	case "c":
		fmt.Println("c!")
		// fallthrough // cannot fallthrough final case in switch
	}

// OUTPUT: 
// go!
// python!
// ruby!
// c!

	switch e := "go"; e {
	case "java":
		fmt.Println("java!")
	case "go":
		fmt.Println("go!")
		fallthrough // 특정 case의 문장을 실행한 뒤 다음 case의 문장을 실행하고 싶을 때는 fallthrough 키워드를 사용합니다. 
	case "golang":
		fmt.Println("golang!")
	case "ruby":
		fmt.Println("ruby!")
	case "c":
		fmt.Println("c!")
		// fallthrough // cannot fallthrough final case in switch
	}

// OUTPUT: 
// go!
// golang!

// fallthrough 사용하기
// 특정 case의 문장을 실행한 뒤 다음 case의 문장을 실행하고 싶을 때는 fallthrough 키워드를 사용합니다. 
// 마치 C, C++의 switch 분기문에서 break 키워드를 생략한 것처럼 동작합니다. 
// 단, 맨 마지막 case에는 fallthrough 키워드를 사용할 수 없습니다.

}

// Go는 case문 마지막에 break 문을 적든 break 문을 생략하든, 
// 항상 break 하여 switch 문을 빠져나온다. 이는 Go 컴파일러가 자동으로 break 문을 각 case문 블럭 마지막에 추가하기 때문이다. 
// Go에서 만약 이러한 디폴트 break 문을 사용하지 않고, C나 C#처럼 계속 밑의 문장들(다음 case문 코드 블럭들)을 실행하게 하려면, 
// fallthrough 문을 명시해 주면 된다. fallthrough 문을 사용한 아래 예제를 실행하면, 
// "2 이하/3 이하/default 도달"을 모두 출력하게 된다. 즉, 일단 case 2 에 도착한 후 fallthrough 가 있으므로, 
// (val 값이 3이 아님에도) case 3의 코드 블럭을 계속 실행하고, case 3에도 fallthrough 가 있으므로 default 블럭을 계속 실행한다.

// package main
 
// import "fmt"
 
// func main() {
//     check(2)
// }
 
// func check(val int) {
//     switch val {
//     case 1:
//         fmt.Println("1 이하")
//         fallthrough
//     case 2:
//         fmt.Println("2 이하")
//         fallthrough
//     case 3:
//         fmt.Println("3 이하")
//         fallthrough
//     default:
//         fmt.Println("default 도달")
//     }
// }


//(2) 특이 사항

  // - fall-through 기능 제어 

  //   > 간혹의 경우 여러개의 case문의 결과를 동일 하게  처리 하기 위해서 case문을 탗출하는 break 문을 

  //     생략 하여 쓰는 경우가 존재합니다. 이렇게 case 문을 제어하는것을 fall-through 라고 합니다.

  //   > typescript에는 이 긴능을 허용 하리 말지를 tsconfig.json파일에  compilerOptions 

  //     의 "noFallthroughCasesInSwitch": true를 통해서 제어 합니다.


// Fall through는 case문에서 break가 빠져서 다음 case문에 있는 로직이 실행되는 경우이다. 아마 초급 개발자라면 이러한 실수를 할수 있다고 생각한다.
// 이와 같은 실수는 tsconfig.json에서 설정을 통해서 방지할 수 있다.
// {
// “compilerOptions” : {
// “noFallthroughCaseInSwitch” : true
// }
// }

// 라고 지정해주면 fall through가 발생하는 경우,
// “Fallingthrough case in Switch” 라는 에러메세지가 나타 난다.

// 물론 위와 같이 설정을 해주어도 개발자 의도적으로 fall through를 발생시킬수는 있다

// let input = 0;

// switch(input){
// case 0:
// case 1:
// console.log(“1”);
// break;
// }