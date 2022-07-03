package main

// 외부함수 closure()
func closure(size int) func() int {
	sub := size         // 첫번째 익명함수에서 sub=0 으로 정의했습니다.
	return func() int { // 내부함수
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
