# Range

## for range
- for 와 range를 함께 사용하여 컬렉션의 요소들을 하나씩 가져와 코드를 반복 실행
- 장점: 슬라이스, 맵, 채널의 크기를 몰라도 개별 원소를 하나씩 처리 가능
- 다른 언어의 foreach 와 비슷함

``` go
// `range expression`으로부터 요소의 인덱스/값을 하나씩 리턴하여 `iteration variables`에 각각 할당한다.
for [`iteration variables`] := range `range expression` {
    ...
}
```

#### Range Expression 별 iteration variable

|Range Expression|1st var|2nd var|
|-------- |-------- |--------|
|array, slice|index|value
|string|index|rune|
|map|key|value
|channel|element|-|
