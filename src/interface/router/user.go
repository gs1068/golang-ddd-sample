package router

import (
	"github.com/gs1068/golang_ddd_sample/interface/handler"
	"github.com/labstack/echo"
)

// InitRouting routesの初期化
func InitUserRouting(e *echo.Echo, userHandler handler.UserHandler) {
	e.POST("/user", userHandler.Post())
	e.GET("/user/:id", userHandler.Get())
	e.PUT("/user/:id", userHandler.Put())
	e.DELETE("/user/:id", userHandler.Delete())
}
