package seed

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
)

type User struct {
	UserName string
}

func UsersSeed(db *gorm.DB) error {
	for i := 0; i < 300; i++ {
		name := "Test User" + strconv.Itoa(i)
		user := User{UserName: name}
		if err := db.Create(&user).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}
