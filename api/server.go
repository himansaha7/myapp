package api

import (
	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()
	setupRouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func setupRouter(e *echo.Echo) {
	e.POST("/employee", saveEmployee)
	e.GET("/employee/:id", getEmployee)
	e.DELETE("/employee/:id", deleteEmployee)
}
