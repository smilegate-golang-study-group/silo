# Variadic-functions
- golang의 가변인자 (variadic arguments) 사용법
- `variadic function`은 가변인자를 받아오는 함수를 의미함.
- 예시: `fmt.Println`
```go
// fmt.Println
func Println(a ...any) (n int, err error)

fmt.Println("println test")
// println test

fmt.Println("from:", 1, 2, 3)
// from: 1 2 3
```

## code
- `append()` - `func append(slice []Type, elems ...Type) []Type`
  - `slice`(=크기가 가변적인 배열, List) 에 여러 개의 원소를 추가할 때
  - 서로 다른 `slice`를 합칠 때
```go
package main

import "fmt"

var array1, array2 []int

func main() {
  //func append(slice []Type, elems ...Type) []Type

  array1 = append(array1, 100)
  fmt.Println("array1:", array1)
  //array1: [100]

  array2 = append(array2, 200, 300)
  array2 = append(array2, []int{400, 500}...)
  fmt.Println("array2:", array2)
  //array2: [200 300 400 500]

  array1 = append(array1, array2...)
  fmt.Println("array1:", array1)
  //array1: [100 200 300 400 500]

}
```
#### 결과:
```bash
array1: [100]
array2: [200 300 400 500]    
array1: [100 200 300 400 500]
```

## struct-version
- 위의 코드와 동일하나 새로 정의한 구조체로 구현.
```go
package main

import "fmt"

type Element struct {
  Variable int
}

var array1, array2 []Element

func main() {
  // func append(slice []Type, elems ...Type) []Type

  array1 = append(array1, Element{Variable: 100})
  fmt.Println("array1:", array1)
  //array1: [{100}]

  array2 = append(array2, Element{Variable: 200}, Element{Variable: 300})
  array2 = append(array2, []Element{{Variable: 400}, {Variable: 500}}...)
  fmt.Println("array2:", array2)
  //array2: [{200} {300} {400} {500}]

  array1 = append(array1, array2...)
  fmt.Println("array1:", array1)
  //array1: [{100} {200} {300} {400} {500}]
}
```
#### 결과:
```bash
array1: [{100}]
array2: [{200} {300} {400} {500}]      
array1: [{100} {200} {300} {400} {500}]
```