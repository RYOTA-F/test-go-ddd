package infra

import (
	"tutorial-go-ddd/domain/model"
	"tutorial-go-ddd/domain/repository"

	"github.com/jinzhu/gorm"
)

type testRepository struct {
	Conn *gorm.DB
}

func NewTestRepotitory(conn *gorm.DB) repository.TestRepository {
	return &testRepository{Conn: conn}
}

func (testRepository *testRepository) FindAll() ([]*model.Test, error) {
	tests := []*model.Test{}

	if err := testRepository.Conn.Find(&tests).Error; err != nil{
		return nil, err
	}

	return tests, nil
}

func (testRepository *testRepository) FindByID(id int) (*model.Test, error) {
	test := &model.Test{Id: id}

	if err := testRepository.Conn.First(&test).Error; err != nil{
		return nil, err
	}

	return test , nil
}
