# Interfaces

## code
```go
package main

import (
	"fmt"
	"math"
)

// interface : method 집합을 그룹화하고 명명하기 위함
type geometry interface { // 인터페이스 정의
	area() float64
	perim() float64
}


// 인터페이스는 rect, circle 유형을 통해 구현함
// 인터페이스 구현 <=> 정의한 모든 method를 구현해야 함
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}


// implementation geometry on "rect"s
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// implementation for "circle"s
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}


// 정의된 인터페이스 내 method 호출하여 사용 가능
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}


func main() {
// geometry 인터페이스 구현은 circle, rect 구조체 타입에 의해 구현됨
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

// 인터페이스가 구현되어 있으므로 instance 사용 가능
	measure(r)
	measure(c)
}
```
#### 결과:
```bash
{3 4}
12
14
{5}
78.53981633974483
31.41592653589793
```