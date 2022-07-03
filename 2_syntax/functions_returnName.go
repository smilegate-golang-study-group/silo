package main

import (
	"fmt"
)

func plus(a, b int) int { //동일한 타입의 매개변수가 연속되면 마지막 매개변수의 타입만 지정해도 된다.

	return a + b
}

// 리턴 값에 이름 붙이기
func plusPlus(a, b, c int) (total int) {
	total = a + b + c
	return total
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

}
