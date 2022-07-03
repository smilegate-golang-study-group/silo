package main

import "fmt"

type Element struct {
	Variable int
}

var array1, array2 []Element

func main() {
	// func append(slice []Type, elems ...Type) []Type

	array1 = append(array1, Element{Variable: 100})
	fmt.Println("array1:", array1)
	//array1: [{100}]

	array2 = append(array2, Element{Variable: 200}, Element{Variable: 300})
	array2 = append(array2, []Element{{Variable: 400}, {Variable: 500}}...)
	fmt.Println("array2:", array2)
	//array2: [{200} {300} {400} {500}]

	array1 = append(array1, array2...)
	fmt.Println("array1:", array1)
	//array1: [{100} {200} {300} {400} {500}]
}
