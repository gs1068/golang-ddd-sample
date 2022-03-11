package repository

import (
    "github.com/gs1068/golang_ddd_sample/domain/model"
)

// TaskRepository task repositoryのinterface
type TaskRepository interface {
    Create(task *model.Task) (*model.Task, error)
    FindByID(id int) (*model.Task, error)
    Update(task *model.Task) (*model.Task, error)
    Delete(task *model.Task) error
}