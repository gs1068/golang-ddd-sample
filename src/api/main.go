package main

import (
	"github.com/gs1068/golang_ddd_sample/config"
	"github.com/gs1068/golang_ddd_sample/infra"
	"github.com/gs1068/golang_ddd_sample/interface/handler"
	"github.com/gs1068/golang_ddd_sample/interface/router"
	"github.com/gs1068/golang_ddd_sample/usecase"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	db := config.NewDB()
	// task
	taskRepository := infra.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUsecase)
	router.InitTaskRouting(e, taskHandler)
	// user
	userRepository := infra.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	router.InitUserRouting(e, userHandler)
	e.Logger.Fatal(e.Start(":8888"))
}
