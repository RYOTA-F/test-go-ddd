package usecase

import (
	"tutorial-go-ddd/domain/model"
	"tutorial-go-ddd/domain/repository"
)

type TestUsecase interface {
	FindAll() ([]*model.Test, error)
	FindByID(id int) (*model.Test, error)
	Create(name string) (*model.Test, error)
}

type testUsecase struct {
	testRepository repository.TestRepository
}

func NewTestUseCase(testRepository repository.TestRepository) TestUsecase {
	return &testUsecase{testRepository: testRepository}
}

func (testUsecase *testUsecase) FindAll() ([]*model.Test, error) {
	tests, err := testUsecase.testRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return tests, nil
}

func (testUsecase *testUsecase) FindByID(id int) (*model.Test, error) {
	test, err := testUsecase.testRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return test, nil
}

func (testUsecase *testUsecase) Create(name string) (*model.Test, error) {
	test, err := model.NewTest(name)

	if err != nil {
		return nil, err
	}

	createdTest, err := testUsecase.testRepository.Create(test)
	if err != nil {
		return nil, err
	}

	return createdTest, nil
}