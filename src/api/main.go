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
	taskRepository := infra.NewTaskRepository(config.NewDB())
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	e := echo.New()
	handler.InitRouting(e, taskHandler)
	e.Logger.Fatal(e.Start(":8888"))
}
