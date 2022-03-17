package usecase

import (
	"github.com/gs1068/golang_ddd_sample/domain/model"
	"github.com/gs1068/golang_ddd_sample/domain/repository"
)

type TaskUsecase interface {
	Create(user_id int, title, content string) (*model.Task, error)
	FindByID(id int) (*model.Task, error)
	Update(id int, user_id int, title, content string) (*model.Task, error)
	Delete(id int) error
}

type taskUsecase struct {
	taskRepo repository.TaskRepository
}

func NewTaskUsecase(taskRepo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: taskRepo}
}

func (tu *taskUsecase) Create(user_id int, title, content string) (*model.Task, error) {
	task, err := model.NewTask(user_id, title, content)
	if err != nil {
		return nil, err
	}

	createdTask, err := tu.taskRepo.Create(task)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (tu *taskUsecase) FindByID(id int) (*model.Task, error) {
	foundTask, err := tu.taskRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundTask, nil
}

func (tu *taskUsecase) Update(id, user_id int, title, content string) (*model.Task, error) {
	targetTask, err := tu.taskRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = targetTask.Set(title, content)
	if err != nil {
		return nil, err
	}

	updatedTask, err := tu.taskRepo.Update(targetTask)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (tu *taskUsecase) Delete(id int) error {
	task, err := tu.taskRepo.FindByID(id)
	if err != nil {
		return err
	}

	err = tu.taskRepo.Delete(task)
	if err != nil {
		return err
	}

	return nil
}
