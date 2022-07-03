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
