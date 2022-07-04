# Structs

## code
```go
package main

import "fmt"

// 구조체 : 유형이 지정된 필드의 모음
// 사용 목적 : 데이터를 그룹화하여 record를 형성하려고 할 때 (생성 및 관리의 유용성)

type person struct { // person 구조체 선언 - name, age 필드로 구성
	name string
	age  int
}

func newPerson(name string) *person { // newPerson 구조체 생성
	p := person{name: name} // 함수 범위의 지역 변수
	p.age = 42
	return &p // 함수 내, 포인터 반환 가능 
}

func main() { // 구조체 초기화
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30}) // 필드 이름 지정 가능
	fmt.Println(person{name: "Fred"}) // 생략된 필드 값 : 0
	fmt.Println(&person{name: "Ann", age: 40}) // 구조체에 대한 포인터 생성 : &
	fmt.Println(newPerson("Jon")) // 생성자를 통해 새 구조체 생성 - 캡슐화

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // 구조체 필드에 엑세스 : .

	sp := &s // 포인터 사용시 자동으로 역참조 됨
	fmt.Println(sp.age)

	sp.age = 51 // 구조체 내부 값 변경 가능
	fmt.Println(sp.age)
}
```
#### 결과:
```bash
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
&{Jon 42}
Sean
50
51
```