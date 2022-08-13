package repository

import "tutorial-go-ddd/domain/model"

type TestRepository interface {
  FindAll() ([]*model.Test, error)
  FindByID(id int) (*model.Test, error)
  Create(test *model.Test) (*model.Test, error)
  Update(test *model.Test) (*model.Test, error)
  Delete(task *model.Test) error
}
