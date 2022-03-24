package repository

import "github.com/gs1068/golang-ddd-sample/domain/model"

type TimelineRepository interface {
	Create(timeline *model.Timeline) (*model.Timeline, error)
	FindByID(id int) (*model.Timeline, error)
	Update(timeline *model.Timeline) (*model.Timeline, error)
	Delete(timeline *model.Timeline) error
}
