package infrastructure

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Init() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, test())
	})
	e.Logger.Fatal(e.Start(":7050"))
}

func test() string {
	fmt.Println("djf")
	return "Test"
}
