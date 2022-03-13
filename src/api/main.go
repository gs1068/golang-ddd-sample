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
	// ①config.NewDB():dbを返している
	// ②dbを含む構造体{Conn: db}をCRUDのインターフェイス型に変換している
	taskRepository := infra.NewTaskRepository(config.NewDB())
	// ③レポジトリのインターフェイスをusecaseのインターフェイスに代入している
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	// ④usecaseのインターフェイスをhandlerのインターフェイスに代入している
	taskHandler := handler.NewTaskHandler(taskUsecase)

	// 上記で、
	// ①infrastructure → ②domain → ③usecase → ④interfaceを実現している
	e := echo.New()
	// ルーティングの定義
	handler.InitRouting(e, taskHandler)
	e.Logger.Fatal(e.Start(":8888"))
}
