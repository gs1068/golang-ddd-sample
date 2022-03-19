package seed

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/jinzhu/gorm"
)

type Timeline struct {
	ID      int
	UserID  int
	Content string
}

func TimelinesSeed(db *gorm.DB, wg *sync.WaitGroup) error {
	defer wg.Done()

	for i := 0; i < 300; i++ {
		contet := "みんなは生涯同時にこの発表者に従って訳の時に歩くんあり。ぼんやり時間で把持ごとは毫もこういう談判ですですだってに聴このにおきでをは創設進んたたいが、当然にも直さたずたです。" + strconv.Itoa(i)
		timeline := Timeline{UserID: i + 1, Content: contet}
		if err := db.Create(&timeline).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}
