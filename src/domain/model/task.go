package model

import (
	"errors"
)

// Task taskの構造体
type Task struct {
	ID      int
	UserID  int
	Title   string
	Content string
	User    User
}

// NewTask taskのコンストラクタ
func NewTask(user_id int, title, content string) (*Task, error) {
	if title == "" {
		return nil, errors.New("titleを入力してください")
	}

	task := &Task{
		UserID:  user_id,
		Title:   title,
		Content: content,
	}

	return task, nil
}

// Set taskのセッター
func (t *Task) Set(title, content string) error {
	if title == "" {
		return errors.New("titleを入力してください")
	}

	t.Title = title
	t.Content = content

	return nil
}
