package main

import "fmt"

func hello(i func(p, q string) string) {
	fmt.Println(i("Hello ", "world"))
}

func welcome() func(i, j string) string {
	myf := func(i, j string) string {
		return i + j + "GoStudy@_@"
	}
	return myf
}

func main() {

	// 익명함수
	func() {
		fmt.Println("Anonymous function")
	}()

	// 익명함수를 변수에 할당
	value := func() {
		fmt.Println("Assigning an anonymous function to a variable")
	}
	value()

	// 익명 함수에 매개변수 전달
	func(ele string) {
		fmt.Println(ele)
	}("Passing arguments")

	// 익명 함수를 다른 함수의 매개변수로 전달
	value2 := func(p, q string) string {
		return p + q + "!!"
	}
	hello(value2)

	// 다른 함수에서 익명 함수를 return
	value3 := welcome()
	fmt.Println(value3("Welcome ", "to "))

}
