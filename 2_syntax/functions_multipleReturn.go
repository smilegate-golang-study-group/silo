package main

import (
	"fmt"
	"io"
	"strings"
)

func sumdiff(a, b int) (sum, diff int) {
	sum = a + b
	diff = a - b
	return sum, diff //return 값을 2개 지정
}

func main() {
	/*
		sumdiff 함수
	*/
	sum, diff := sumdiff(5, 6)
	fmt.Println(sum, diff)

	// return 값 중 사용하지 않는 값은 '_'로 할당
	_, onlyDiff := sumdiff(5, 6)
	fmt.Println(onlyDiff)

	/*
		io.ReadFull() 함수
		func ReadFull(r Reader, buf []byte) (n int, err error)
	*/

	// Defining reader using NewReader method
	reader := strings.NewReader("Geeks")

	// Defining buffer of specified length using make keyword
	buffer := make([]byte, 4)

	// Calling ReadFull method with its parameters
	n, err := io.ReadFull(reader, buffer)

	// If error is not nil then panics
	if err != nil {
		panic(err)
	}

	// Prints output
	fmt.Printf("Number of bytes in the buffer: %d\n", n)
	fmt.Printf("Content in buffer: %s\n", buffer)

	/*
		버퍼의 길이를 reader 보다 크게하여 EOF 오류 발생 확인.
	*/
	// Defining reader using NewReader method
	reader2 := strings.NewReader("Geeks")

	// Defining buffer of specified length using make keyword
	buffer2 := make([]byte, 6)

	// Calling ReadFull method with its parameters
	n2, err2 := io.ReadFull(reader2, buffer2)

	// If error is not nil then panics
	if err2 != nil {
		panic(err2)
	}

	// Prints output
	fmt.Printf("Number of bytes in the buffer: %d\n", n2)
	fmt.Printf("Content in buffer: %s\n", buffer2)
}
