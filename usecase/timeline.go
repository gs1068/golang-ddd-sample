package usecase

import (
	"github.com/gs1068/golang_ddd_sample/domain/model"
	"github.com/gs1068/golang_ddd_sample/domain/repository"
)

type TimelineUsecase interface {
	Create(user_id int, content string) (*model.Timeline, error)
	FindByID(id int) (*model.Timeline, error)
	Update(id int, user_id int, content string) (*model.Timeline, error)
	Delete(id int) error
}

type timelineUsecase struct {
	timelineRepo repository.TimelineRepository
}

func NewTimelineUsecase(timelineRepo repository.TimelineRepository) TimelineUsecase {
	return &timelineUsecase{timelineRepo: timelineRepo}
}

func (tu *timelineUsecase) Create(user_id int, content string) (*model.Timeline, error) {
	timeline, err := model.NewTimeline(user_id, content)
	if err != nil {
		return nil, err
	}

	createdTimeline, err := tu.timelineRepo.Create(timeline)
	if err != nil {
		return nil, err
	}

	return createdTimeline, nil
}

func (tu *timelineUsecase) FindByID(id int) (*model.Timeline, error) {
	foundTimeline, err := tu.timelineRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundTimeline, nil
}

func (tu *timelineUsecase) Update(id, user_id int, content string) (*model.Timeline, error) {
	targetTimeline, err := tu.timelineRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = targetTimeline.Set(content)
	if err != nil {
		return nil, err
	}

	updatedTimeline, err := tu.timelineRepo.Update(targetTimeline)
	if err != nil {
		return nil, err
	}

	return updatedTimeline, nil
}

func (tu *timelineUsecase) Delete(id int) error {
	timeline, err := tu.timelineRepo.FindByID(id)
	if err != nil {
		return err
	}

	err = tu.timelineRepo.Delete(timeline)
	if err != nil {
		return err
	}

	return nil
}
