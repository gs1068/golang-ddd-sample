package main

import (
	"github.com/gs1068/golang-ddd-sample/config"
	"github.com/gs1068/golang-ddd-sample/infra"
	"github.com/gs1068/golang-ddd-sample/interface/handler"
	"github.com/gs1068/golang-ddd-sample/interface/router"
	"github.com/gs1068/golang-ddd-sample/usecase"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

// main
func main() {
	e := echo.New()
	db := config.NewDB()
	// user
	userRepository := infra.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	router.InitUserRouting(e, userHandler)
	// task
	taskRepository := infra.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUsecase)
	router.InitTaskRouting(e, taskHandler)
	// timeline
	timelineRepository := infra.NewTimelineRepository(db)
	timelineUsecase := usecase.NewTimelineUsecase(timelineRepository)
	timelineHandler := handler.NewTimelineHandler(timelineUsecase)
	router.InitTimelineRouting(e, timelineHandler)
	e.Logger.Fatal(e.Start(":8888"))
}
