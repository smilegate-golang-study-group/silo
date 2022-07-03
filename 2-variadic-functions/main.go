package main

import "fmt"

var array1, array2 []int

func main() {
	//func append(slice []Type, elems ...Type) []Type

	array1 = append(array1, 100)
	fmt.Println("array1:", array1)
	//array1: [100]

	array2 = append(array2, 200, 300)
	array2 = append(array2, []int{400, 500}...)
	fmt.Println("array2:", array2)
	//array2: [200 300 400 500]

	array1 = append(array1, array2...)
	fmt.Println("array1:", array1)
	//array1: [100 200 300 400 500]

}
