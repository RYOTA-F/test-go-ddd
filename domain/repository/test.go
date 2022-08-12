package repository

import "tutorial-go-ddd/domain/model"

type TestRepository interface {
  FindAll() ([]*model.Test, error)
  FindByID(id int) (*model.Test, error)
}
