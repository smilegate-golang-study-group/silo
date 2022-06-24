package main

import (
	"fmt"

	// 여기에 외부 패키지를 추가할 수 있습니다.
	"go.uber.org/zap"
)

func main() {
	fmt.Println("Hello, World!")

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Info("zap test OK")
}
