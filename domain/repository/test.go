package repository

import "tutorial-go-ddd/domain/model"

type TestRepository interface {
  Search(name string) ([]*model.Test, error)
}
