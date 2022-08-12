package handler

import (
	"net/http"
	"strconv"
	"tutorial-go-ddd/usecase"

	"github.com/labstack/echo"
)

type TestHandler interface {
	Get() echo.HandlerFunc
}

type testHandler struct {
	testUsecase usecase.TestUsecase
}

func NewTestHandler(testUsecase usecase.TestUsecase) TestHandler {
	return &testHandler{testUsecase: testUsecase}
}

type requestTest struct {
	Name string `json:"name"`
}

type responseTest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (testHandler *testHandler) Get() echo.HandlerFunc {
	return func(context echo.Context) error {
		id, err := strconv.Atoi((context.Param("id")))
		if err != nil {
			return context.JSON(http.StatusBadRequest, err.Error())
		}

		foundTest, err := testHandler.testUsecase.FindByID(id)
		if err != nil {
			return context.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTest{
			Id: foundTest.Id,
			Name: foundTest.Name,
		}

		return context.JSON(http.StatusOK, res)
	}
}