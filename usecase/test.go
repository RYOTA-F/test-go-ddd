package usecase

import (
	"tutorial-go-ddd/domain/model"
	"tutorial-go-ddd/domain/repository"
)

type TestUseCase interface {
	Search(name string) ([]*model.Test, error)
}

type testUseCase struct {
	testRepository repository.TestRepository
}

func NewTestUseCase(tr repository.TestRepository) TestUseCase {
	return &testUseCase{
		testRepository: tr,
	}
}

func (tu testUseCase) Search(name string) (tests []*model.Test, err error) {
	tests, err = tu.testRepository.Search(name)
	if err != nil {
		return nil, err
	}

	return tests, nil
}