# Function

- 여러 문장을 묶은 코드 블럭
- 코드 재사용을 통한 과도한 메모리 사용을 줄이고 가독성 향상
- 패키지 안에 정의되며 호출되는 함수가 호출하는 함수보다 앞에 위치하지 않아도 됨
- *multiple return value* 가능

``` go
/*
선언: func 키워드로 선언. 선언 위치는 어느 곳이든 가능
func_name: 함수 이름
parameter_list: 매개변수 이름, 타입
return_type: 반환 값의 타입. 값 이름 지정 가능. return_type 지정 시 함수에 return 문을 포함해야 함.
*/
func func_name(parameter_list)(return_type) //반환값 존재
func func_name(parameter_list)              //매개변수 존재, 반환 값 없음
func func_name()(return_type)               //매개변수 없음, 반환 값 존재
func func_name()                            //매개변수 없음, 반환 값 없음
```




## Function Argument
- actual parameter: 함수에 실제로 전달되는 파라미터
- formal parameter: 함수가 전달받은 파라미터

Go에서 파라미터를 전달하는 방식은 크게 Pass By Value와 Pass By Reference로 나뉜다.
- ### Call by value
    - actual parameter의 값이 formal parameter로 복사됨.
    - 서로 다른 메모리 영역에 저장된다.
    - 함수 안에서 발생하는 변경사항은 actual parameter에 영향을 주지 않음.
- ### Call by reference
    - actual parameter와 formal parameter가 동일한 메모리 영역을 참조
    - 함수 안에서 발생하는 변경사항은 actual parameter에도 적용됨.

- callBy.go

## 함수 유형
### main() 함수
- go는 패키지 이름이 main이고 그 안에 main() 함수만 정의되어 있으면 독립 실행 프로그램으로 취급
- 소스 파일 이름에 관계 없이 main() 함수가 있는 곳에서 프로그램을 실행
- 한 프로젝트 안에서 main()이 정의된 파일을 하나여야 함


### 익명 함수
```go
func(parameter_list)(return_type){
// code..
return // return_type 이 있으면 return문 사용
}()
```
- 이름 없이 정의되는 함수(function literal이라고도 함.)
- 로컬에서 간단한 작업을 수행할 때만 사용하는 것이 좋다.
- 변수에 익명 함수를 할당할 수 있다. 변수 타입은 함수.
- 익명 함수를 다른 함수의 매개변수로 전달할 수 있다.
- 다른 함수에서 익명 함수를 리턴할 수 있다.
- anonymousFunc.go


### 여러 값을 리턴하는 함수
- 장점: 함수가 리턴하는 여러 값에 접근하기 위한 구조체를 정의하지 않아도 된다.
- 리턴 값과 함수의 리턴값을 할당받는 변수의 개수가 일치하지 않으면 에러 발생
- 리턴 값 중 일부만 사용하고 싶다면 사용하지 않을 리턴값은 '_' 로 처리
- 활용: 결과와 오류를 모두 반환하기 위해 관용적으로 사용됨.
- multipleReturn.go

### 리턴 값에 이름 붙이기
- 함수의 리턴 값에 이름을 붙일 수 있다.
- return 문 쓰지 않으면 에러 발생
- 필수는 아니지만 코드 가독성을 높이고 길이를 줄일 수 있다.
- 함수가 시작될 때 각 타입의 zero value로 초기화되고 함수의 return문이 실행될 때 result parameter 값이 retrun value로 사용 
- return문에 아무 인수도 지정하지 않으면 각 리턴 값이 선언된 순서대로 리턴함
- returnName.go

### 포인터를 매개변수로 받는 함수
### 포인터를 리턴하는 함수
### 함수를 리턴하는 함수
### 함수를 매개변수로 받는 함수
