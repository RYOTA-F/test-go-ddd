package infra

import (
	"tutorial-go-ddd/domain/model"
	"tutorial-go-ddd/domain/repository"

	"github.com/jinzhu/gorm"
)

type testRepository struct {
	Connect *gorm.DB
}

func NewTestRepotitory(conn *gorm.DB) repository.TestRepository {
	return &testRepository{Connect: conn}
}

func (testRepository *testRepository) FindAll() ([]*model.Test, error) {
	tests := []*model.Test{}

	if err := testRepository.Connect.Find(&tests).Error; err != nil{
		return nil, err
	}

	return tests, nil
}

func (testRepository *testRepository) FindByID(id int) (*model.Test, error) {
	test := &model.Test{Id: id}

	if err := testRepository.Connect.First(&test).Error; err != nil{
		return nil, err
	}

	return test , nil
}

func (testRepository *testRepository) Create(test *model.Test) (*model.Test, error) {
	if err := testRepository.Connect.Create(&test).Error; err != nil {
		return nil, err
	}

	return test, nil
}

func (testRepository *testRepository) Update(test *model.Test) (*model.Test, error) {
	if err := testRepository.Connect.Model(&test).Update(&test).Error; err != nil {
		return nil, err
	}

	return test, nil
}