# Error

## Go의 Error 처리
- Go는 다른 객체 지향 언어에서 사용하는 try/catch 문을 제공하지 않는다.
- C 처럼 에러가 발생한 지점에서 에러를 처리한다.
- error type: Go는 error 라는 내장된 interface 타입을 제공
- Go 언어에서도 에러 메시지를 제공하지만 에러 타입을 직접 정의할 수도 있다.
- Go는 관례 상 함수의 리턴 값 중 가장 마지막 값을 에러로 반환(multiple return value)
- 함수가 결과와 에러를 리턴한다면, 에러가 nil 인지 체크하여 에러를 처리할 수 있다.
- error의 Type을 체크해서 에러 타입별로 별도의 에러 처리도 가능함.
``` go
// 자주 사용하는 에러 처리 방식
if err != nil {
    fmt.println(err) // 에러를 화면에 출력
    log.println(err) // 에러를 로그 서비스로 보냄
    panic(err)       // 프로그램을 즉시 종료
   // os.Exit(10) 프로그램 종료
    }
```



## Custom Error
### error interface
- 에러 정보를 담은 string을 반환
- Error() method 구현만으로 모든 코드에서 사용 가능
- Error() method를 가지면 어떤 타입이든 error로 인식된다.
``` go
type error interface {
    Error() string
}
```

### errorString
- errors 패키지에 정의된 구조체
- Error() method 를 이용해 s 값을 가져올 수 있다. 
``` go
type errorString struct {
	s string
}
```

### Error 값 생성
- errors 패키지의 New 함수로 errorString 포인터 생성
``` go
import "errors"
var ErrInvalidParam = errors.New("Mypackage: invalid parameter")
```
- New 함수 호출하면 errorString에 값을 설정하고 포인터를 반환
``` go
func New(text string) error {
	return &errorString(text)
}
```
- fmt.Errorof 로 형식을 지정하여 에러 생성
``` go
func Errorf(format string, a ...interface{}) error
/*
parameter
string: %s, %d 등 placeholder를 포함한 에러 메시지
a ...interface{}: 코드나 내장함수에 사용되는 constant variable name 
*/
var ErrInvalidParam = fmt.Errorf("Invalid parameter [%s]", param)
```
