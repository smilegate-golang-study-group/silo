# Switch

## code
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i { // 기본 형태
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() { // 쉼표를 통해 다중 조건을 한 명령문에서 분기
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default: // 기본값도 설정 가능
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch { // if-else 구문처럼 사용하는 형태
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) { // 값 대신 유형을 비교하는 방법
		switch t := i.(type) { // 인터페이스 값의 유형을 비교
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
```
#### 결과:
```bash
Write 2 as two
It's a weekday
It's before noon
I'm a bool
I'm an int
Don't know type string
```