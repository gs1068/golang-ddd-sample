package config

import (
	"github.com/gs1068/golang_ddd_sample/domain/model"

	"github.com/jinzhu/gorm"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(db)/dddsample")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.Task{})

	return db
}
