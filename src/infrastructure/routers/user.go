package routers

import (
	"github.com/labstack/echo/v4"
	"go_todo/infrastructure"
	"go_todo/interface/controllers"
)

func UserRouter(e *echo.Echo) {
	userController := controllers.NewUserController(infrastructure.NewSqlHandler())
	e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	e.GET("/user/:id", func(c echo.Context) error { return userController.Show(c) })
	e.POST("/create", func(c echo.Context) error { return userController.Create(c) })
	e.PUT("/user/:id", func(c echo.Context) error { return userController.Save(c) })
	e.DELETE("/user/:id", func(c echo.Context) error { return userController.Delete(c) })
}
