# If_Else

## code
```go
package main

import "fmt"

func main() {

	if 7/2 == 0 { // if-else 구문
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 { // 단순 if문
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 { // 명령문에 사용된 num은 모든 분기에서 사용가능
		fmt.Println(num, "is negative")
	} else if num < 10 { // 조건 주위에 괄호 생략 가능, 중괄호는 필요
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
```
#### 결과:
```bash
7 is odd
8 is divisible by 4
9 has 1 digit
```