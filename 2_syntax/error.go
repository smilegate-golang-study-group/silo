package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	/*
		os.Open() 함수: File 포인터와 error 인터페이스를 리턴한다.
		func Open(name string) (file *File, err error)
		log.Fatal() 함수: 메시지 출력 후 os.Exit(1)을 호출하여 프로그램을 종료한다.

		아래 코드는 파일 열기에 실패하면 에러를 출력하고 종료한다.

		2022/07/03 19:49:01 open C:\golang\src\5_error\test.txt: The system cannot find the file specified.
		exit status 1

	*/
	f, err := os.Open("C:\\golang\\src\\5_error\\test.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	println(f.Name())

	/*
		errors.New() 함수로 custom error 생성
		정사각형 면적을 출력
		에러 체크 - 음수인 input이 들어오면 에러를 출력하기

	*/
	side := 2.0
	err, area := findAreaSquare(side)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Printf("Area= %f", area)

	/*
		fmt.Errorf() 함수로 custom error 생성
		input parameter가 1이면 "Hi 1" 출력
		에러 체크 - input parameter가 1이 아니면 에러 발생(에러 발생 시간 출력하기)
	*/
	hiS, hiErr := hello(1)
	if hiErr != nil {
		log.Fatal(hiErr)
	}
	fmt.Println(hiS) // Hi 1

	/*
		에러 구조체 사용
	*/
	helloS, helloErr := hello(2)
	if helloErr != nil {
		log.Fatal(helloErr)
	}
	fmt.Println(helloS)

}

// errors.New() 함수로 custom error 생성
func findAreaSquare(side float64) (error, float64) {
	if side < 0.0 {
		return errors.New("Negative value entered!"), 0
	}
	return nil, side * side
}

// fmt.Errorf() 함수로 custom error 생성
func hi(n int) (string, error) {
	if n == 1 {
		s := fmt.Sprintln("Hi ", n)
		return s, nil //정상 동작 - 에러는 nil
	}
	return "", fmt.Errorf("%d는 1이 아님. error occurred at: %v", n, time.Now())
}

// 에러 타입 생성
type HelloError struct {
	time  time.Time
	value int
}

func (e HelloError) Error() string { // 에러 메시지를 생성하여 리턴하는 Error 함수 구현
	return fmt.Sprintf("%v: %d는 1이 아닙니다.", e.time, e.value)
}

func hello(n int) (string, error) {
	if n == 1 { // 1일때만
		s := fmt.Sprintln("Hello ", n) // Hello 문자열을 리턴
		return s, nil                  // 정상 동작이므로 에러 값은 nil
	}
	return "", HelloError{time.Now(), n} // 1이 아니면 에러 구조체 인스턴스를 생성하여 리턴
}
