package repository

import "tutorial-go-ddd/domain/model"

type TestRepository interface {
  FindByID(id int) (*model.Test, error)
}
