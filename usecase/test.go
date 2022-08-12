package usecase

import (
	"tutorial-go-ddd/domain/model"
	"tutorial-go-ddd/domain/repository"
)

type TestUsecase interface {
	FindByID(id int) (*model.Test, error)
}

type testUsecase struct {
	testRepository repository.TestRepository
}

func NewTestUseCase(testRepository repository.TestRepository) TestUsecase {
	return &testUsecase{testRepository: testRepository}
}

func (testUsecase *testUsecase) FindByID(id int) (*model.Test, error) {
	foundTest, err := testUsecase.testRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundTest, nil
}