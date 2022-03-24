package router

import (
	"github.com/gs1068/golang-ddd-sample/interface/handler"
	"github.com/labstack/echo"
)

func InitTimelineRouting(e *echo.Echo, timelineHandler handler.TimelineHandler) {
	e.POST("/timeline", timelineHandler.Post())
	e.GET("/timeline/:id", timelineHandler.Get())
	e.PUT("/timeline/:id", timelineHandler.Put())
	e.DELETE("/timeline/:id", timelineHandler.Delete())
}
