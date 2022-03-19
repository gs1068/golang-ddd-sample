package router

import (
	"github.com/gs1068/golang_ddd_sample/interface/handler"
	"github.com/labstack/echo"
)

// InitRouting routesの初期化
func InitTaskRouting(e *echo.Echo, taskHandler handler.TaskHandler) {
	e.POST("/task", taskHandler.Post())
	e.GET("/task/:id", taskHandler.Get())
	e.GET("/task/pl/:id", taskHandler.GetPL())
	e.PUT("/task/:id", taskHandler.Put())
	e.DELETE("/task/:id", taskHandler.Delete())
}
