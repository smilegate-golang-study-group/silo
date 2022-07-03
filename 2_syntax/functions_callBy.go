package main

import "fmt"

func swap_val(a, b int) int {
	var o int
	o = a
	a = b
	b = o
	return o
}

func swap_ref(a, b *int) int {
	var o int
	o = *a
	*a = *b
	*b = o
	return o
}

func main() {
	var p int = 10
	var q int = 20
	fmt.Printf("p = %d and q = %d\n", p, q)

	// call by values
	swap_val(p, q)
	fmt.Printf("p = %d and q = %d\n", p, q)

	var a int = 10
	var b int = 20
	fmt.Printf("a = %d and b = %d\n", a, b)

	//call by reference
	swap_ref(&a, &b)
	fmt.Printf("a = %d and b = %d", a, b)

}
