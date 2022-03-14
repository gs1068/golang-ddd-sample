package main

import (
	"github.com/gs1068/golang_ddd_sample/config"
	"github.com/gs1068/golang_ddd_sample/infra"
	"github.com/gs1068/golang_ddd_sample/interface/handler"
	"github.com/gs1068/golang_ddd_sample/usecase"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// task
	taskRepository := infra.NewTaskRepository(config.NewDB())
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUsecase)
	e.POST("/task", taskHandler.Post())
	e.GET("/task/:id", taskHandler.Get())
	e.PUT("/task/:id", taskHandler.Put())
	e.DELETE("/task/:id", taskHandler.Delete())

	// user
	userRepository := infra.NewUserRepository(config.NewDB())
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	e.POST("/user", userHandler.Post())
	e.GET("/user/:id", userHandler.Get())
	e.PUT("/user/:id", userHandler.Put())
	e.DELETE("/user/:id", userHandler.Delete())
	e.Logger.Fatal(e.Start(":8888"))
}
