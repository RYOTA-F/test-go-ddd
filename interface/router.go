package handler

import (
	"github.com/labstack/echo"
)


func InitRouting(e *echo.Echo, testHandler TestHandler) {
	e.GET("/tests", testHandler.Index())
	e.GET("/tests/:id", testHandler.Get())
}