package infra

import (
	"github.com/gs1068/golang_ddd_sample/domain/model"
	"github.com/gs1068/golang_ddd_sample/domain/repository"
	"github.com/jinzhu/gorm"
)

type TimelineRepository struct {
	Conn *gorm.DB
}

func NewTimelineRepository(conn *gorm.DB) repository.TimelineRepository {
	return &TimelineRepository{Conn: conn}
}

func (tr *TimelineRepository) Create(timeline *model.Timeline) (*model.Timeline, error) {
	if err := tr.Conn.Create(&timeline).Error; err != nil {
		return nil, err
	}

	return timeline, nil
}

func (tr *TimelineRepository) FindByID(id int) (*model.Timeline, error) {
	timeline := &model.Timeline{ID: id}

	if err := tr.Conn.First(&timeline).Error; err != nil {
		return nil, err
	}

	return timeline, nil
}

func (tr *TimelineRepository) Update(timeline *model.Timeline) (*model.Timeline, error) {
	if err := tr.Conn.Model(&timeline).Update(&timeline).Error; err != nil {
		return nil, err
	}

	return timeline, nil
}

func (tr *TimelineRepository) Delete(timeline *model.Timeline) error {
	if err := tr.Conn.Delete(&timeline).Error; err != nil {
		return err
	}

	return nil
}
