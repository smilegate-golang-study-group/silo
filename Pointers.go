package main

import "fmt"

func zeroval(a int) {
	a = 30
}

func zeroptr(b *int) {
	*b = 3
}

func square(i int) int {

	return i * i
}

func main() {

	var i int = 100
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)
	fmt.Println("zeroptr:", i)

	square(i)
	fmt.Println("square:", square(i))

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

}
