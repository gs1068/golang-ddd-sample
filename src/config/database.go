package config

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(db)/dddsample")
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	return db
}
