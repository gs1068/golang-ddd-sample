package model

import "errors"

type Timeline struct {
	ID      int
	UserID  int
	Content string
}

func NewTimeline(user_id int, content string) (*Timeline, error) {
	if content == "" {
		return nil, errors.New("contentを入力してください")
	}

	timeline := &Timeline{
		UserID:  user_id,
		Content: content,
	}

	return timeline, nil
}

func (t *Timeline) Set(content string) error {
	if content == "" {
		return errors.New("contentを入力してください")
	}

	t.Content = content

	return nil
}
