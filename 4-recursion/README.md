# Recursion
- 재귀함수 사용이 가능합니다.
- 다만 재귀 자체가 시간복잡도 문제가 있어서...

## code
```go
package main

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {

	println(factorial(7))

}
```
#### 결과:
```bash
5040
```