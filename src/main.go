package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, test())
	})
	e.Logger.Fatal(e.Start(":7050"))
}

func test() string {
	fmt.Println("akdsjf")
	return "Test"
}
