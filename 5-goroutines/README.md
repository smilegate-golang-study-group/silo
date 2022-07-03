# Goroutines
- golang에서 제공하는 경량 스레드 입니다.
- 함수를 비동기적으로 동시에 실행할 때 사용합니다.

## code
- `defer`: 함수의 return 직전에 수행되며, clean-up 작업 등에 주로 사용.
```go
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
```
#### 결과:
- 위의 코드는 실행할 때마다 결과가 달라집니다.
```bash
from:  5
from:  1
from:  4
from:  3
from:  2
```

## multi port server

```go
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func start(port string) {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.DEBUG)

	e.GET("/*", func(c echo.Context) error {
		c.Logger().Debugf("ok - server: %s", port)
		return c.String(http.StatusOK, "Hello from "+port)
	})

	e.Logger.Fatal(e.Start(":" + port))
	defer wg.Done()
}

func main() {
	wg.Add(2)

	go start("8081")
	go start("8082")

	wg.Wait()
}
```

#### 결과:
```bash
# 1. go run main.go
⇨ http server started on [::]:8081
⇨ http server started on [::]:8082

# 2. 새로운 터미널창에서 호출 테스트
❯ curl localhost:8082
Hello from 8082% # 응답
❯ curl localhost:8081
Hello from 8081% # 응답

# 3. 호출하면 서버에서 로그를 출력합니다.
{"time":"2022-07-03T21:47:55.070093+09:00","level":"DEBUG","prefix":"echo","file":"main.go","line":"18","message":"ok - server: 8082"}
{"time":"2022-07-03T21:47:56.9486207+09:00","level":"DEBUG","prefix":"echo","file":"main.go","line":"18","message":"ok - server: 8081"}
```