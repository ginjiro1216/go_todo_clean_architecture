package routers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// user系のrouterを注入
	UserRouter(e)
	e.Logger.Fatal(e.Start(":7050"))
}

func test() string {
	fmt.Println("djf")
	return "Test"
}
