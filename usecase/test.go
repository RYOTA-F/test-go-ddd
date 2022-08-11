package usecase

import (
	"tutorial-go-ddd/domain/model"
	"tutorial-go-ddd/domain/repository"
)

type TestUsecase interface {
	Search(name string) ([]*model.Test, error)
}

type testUsecase struct {
	testRepository repository.TestRepository
}

func NewTestUseCase(testRepository repository.TestRepository) TestUsecase {
	return &testUsecase{
		testRepository: testRepository,
	}
}

func (testUsecase testUsecase) Search(name string) (tests []*model.Test, err error) {
	tests, err = testUsecase.testRepository.Search(name)
	if err != nil {
		return nil, err
	}

	return tests, nil
}