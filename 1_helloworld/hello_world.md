## 'hello_world'

사전 설정
```sh
터미널>go mod init example/hello
go: creating new go.mod: module example/hello
go: to add module requirements and sums:
        go mod tidy
```

소스코드

```sh
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

출력

```sh
터미널>go run .
Hello, World!
```
