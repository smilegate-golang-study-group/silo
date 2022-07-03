# For

## code
```go
package main

import "fmt"

func main() {
	i := 1
	for i <= 3 { // 하나의 조건
		fmt.Println(i)
		i += 1
	}

	for j := 7; j <= 9; j++ { // 초기화, 조건, 다음동작
		fmt.Println(j)
	}

	for { // break문을 만나기 전까지 loop를 수행함
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ { // 다음 반복으로 이동
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
```
#### 결과:
```bash
1
2   
3   
7   
8   
9   
loop
1
3
5
```