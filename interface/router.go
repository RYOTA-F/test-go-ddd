package handler

import (
	"github.com/labstack/echo"
)


func InitRouting(e *echo.Echo, testHandler TestHandler) {
	e.GET("/tests", testHandler.Index())
	e.GET("/tests/:id", testHandler.Get())
	e.POST("/tests", testHandler.Post())
	e.PUT("/tests/:id", testHandler.Put())
	e.DELETE("/tests/:id", testHandler.Delete())
}