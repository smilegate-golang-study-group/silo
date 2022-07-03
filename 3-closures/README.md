# Closures
- 함수가 종료되어도, 함수 안의 변수를 계속 사용할 수 있습니다.
- closure로 구현하면, 함수 바깥의 변수를 사용할 수 있습니다. (전역 변수를 대신하는 효과)
- 익명 함수(anonymous function)를 통해 구현할 수 있습니다.
- 참고: https://hwan-shell.tistory.com/339

## code
```go
package main

func closure(size int) func() int {
	sub := size // 첫번째 익명함수에서 sub=0 으로 정의했습니다.
	return func() int {
		sub++ // 두번째 익명함수에서 첫번째 익명함수의 변수인 sub 를 가져올 수 있습니다.
		return sub
	}
}

func main() {
	// num1 을 closure(0)으로 정의합니다.
	num1 := closure(0)
	println(num1()) // 1
	println(num1()) // 2
	println(num1()) // 3

	// num2 를 closure(10)으로 정의합니다.
	num2 := closure(10)
	println(num2()) // 11
	println(num2()) // 12

	println(num1()) // 4
}
```
#### 결과:
```bash
1
2
3
11
12
4
```
#### 디버깅:
```bash
❯ go run -gcflags '-m' main.go
# command-line-arguments                     
./main.go:4:6: can inline closure            
./main.go:6:9: can inline closure.func1      
./main.go:14:17: inlining call to closure    
./main.go:6:9: can inline main.func1         
./main.go:15:14: inlining call to main.func1 
./main.go:16:14: inlining call to main.func1 
./main.go:17:14: inlining call to main.func1 
./main.go:20:17: inlining call to closure    
./main.go:6:9: can inline main.func2         
./main.go:21:14: inlining call to main.func2 
./main.go:22:14: inlining call to main.func2 
./main.go:24:14: inlining call to main.func1 
./main.go:5:2: moved to heap: sub            
./main.go:6:9: func literal escapes to heap  
./main.go:14:17: func literal does not escape
./main.go:20:17: func literal does not escape
1
2
3
11
12
4
```


## 익명 함수(anonymous function)
```go
package main

func sum(a int, b int) int {
	a *= 10
	b *= 10
	return a * b
}

func main() {
	a := 2
	b := 3

	// sum 이라는 함수를 정의해서 사용
	println(sum(a, b))

	// sum과 동일한 기능을 익명함수로 작성
	println(
		func() int {
			a *= 10
			b *= 10
			return a * b
		}())

	// sum과 동일한 기능을 익명함수로 작성 - 2
	println(
		func(a int, b int) int {
			a *= 10
			b *= 10
			return a * b
		}(2, 3))

}
```
#### 결과:
```bash
600
600
600
```