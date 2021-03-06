package seed

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/jinzhu/gorm"
)

type Task struct {
	ID      int
	UserID  int
	Title   string
	Content string
}

func TasksSeed(db *gorm.DB, wg *sync.WaitGroup) error {
	defer wg.Done()
	for i := 0; i < 300; i++ {
		text := "タスク" + strconv.Itoa(i)
		contet := "みんなは生涯同時にこの発表者に従って訳の時に歩くんあり。ぼんやり時間で把持ごとは毫もこういう談判ですですだってに聴このにおきでをは創設進んたたいが、当然にも直さたずたです。" + strconv.Itoa(i)
		task := Task{UserID: i + 1, Title: text, Content: contet}
		if err := db.Create(&task).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}
