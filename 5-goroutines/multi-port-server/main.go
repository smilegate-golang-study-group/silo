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
