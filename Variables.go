package main

import "fmt"

/*
Variables
1) var a = "initial"
fmt.Println(a)
-> 문자열 변수 출력

2) var b, c int = 1, 2
    fmt.Println(b, c)
-> 2개 정수형 변수 선언후 출력

3) var d = true
   fmt.Println(d)
-> true 선언후 출력

4) var e int
   fmt.Println(e)
-> 정수형 디폴트값 출력
*/

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

}
