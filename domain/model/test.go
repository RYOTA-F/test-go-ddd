package model

import "errors"

type Test struct {
  Id int
  Name string
}

func NewTest (name string) (*Test, error) {
  if name == "" {
		return nil, errors.New("名前を入力してください")
	}

  test := &Test{
		Name:   name,
	}

	return test, nil
}