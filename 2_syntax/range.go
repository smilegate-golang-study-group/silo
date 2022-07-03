package main

import "fmt"

func main() {
	var arr [5]string = [5]string{"a", "b", "c", "d", "e"}

	// array의 인덱스, 값 출력
	for idx, str := range arr {
		fmt.Printf("arr[%d]의 값은 %s입니다.\n", idx, str)
	}

	// array의 요소값만 출력
	for _, str := range arr {
		fmt.Printf("%s\n", str)
	}

	// array의 인덱스만 출력
	for idx := range arr {
		fmt.Printf("%d\n", idx)
	}

	// string의 인덱스, rune 출력
	for i, c := range "HelloWorld!" {
		fmt.Println(i, c, string(c))
	}

	// map의 key, value 출력
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// map의 key 출력
	for k := range kvs {
		fmt.Println("key:", k)
	}

	nums := []int{2, 3, 4}
	//array 요소의 합계
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	// 특정 요소값의 인덱스 구하기
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
}
