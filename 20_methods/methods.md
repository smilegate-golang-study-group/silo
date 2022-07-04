# Methods

## code
```go
package main

import "fmt"

type rect struct {
	width, height int
}


// method 정의 2가지 방법으로 가능 - receiver가 포인터 or 값
func (r *rect) area() int { // receiver type = *rect
	return r.width * r.height
}

func (r rect) perim() int { // receiver type = rect
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

// 메서드 호출에 대한 값과 포인터 간의 변환을 자동으로 처리
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}
```
#### 결과:
```bash
area:  50
perim:  30
area:  50
perim:  30
```