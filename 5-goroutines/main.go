package main

import "sync"

var wait sync.WaitGroup

func echo(num int) {
	println("from: ", num)
	defer wait.Done() // 3. 고루틴용 함수가 끝날 때 Done()을 호출
}

func main() {
	num := 5
	wait.Add(num) // 1. 실행할 고루틴의 개수를 지정합니다.

	for i := 1; i <= num; i++ {
		//wait.Add(1) // wait 개수 - 이렇게 정의할 수도 있음.
		go echo(i) // 2. echo 라는 함수가 비동기적으로 실행됩니다.
	}

	// time.Sleep(time.Second * 3)
	// 이렇게 몆초 기다려도 확인할 수 있지만... waitGroup 을 사용하는게 더 일반적

	wait.Wait() // 4. 모든 고루틴이 끝날 때까지 대기합니다.
}
