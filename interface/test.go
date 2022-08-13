package handler

import (
	"net/http"
	"strconv"
	"tutorial-go-ddd/usecase"

	"github.com/labstack/echo"
)

type TestHandler interface {
	Index() echo.HandlerFunc
	Get() echo.HandlerFunc
	Post() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type testHandler struct {
	testUsecase usecase.TestUsecase
}

func NewTestHandler(testUsecase usecase.TestUsecase) TestHandler {
	return &testHandler{testUsecase: testUsecase}
}

type requestTest struct {
	Name	string `json:"name"`
}

type responseTest struct {
	Id   			int    		`json:"id"`
	Name 			string		`json:"name"`
}

func (testHandler *testHandler) Index() echo.HandlerFunc {
	return func(context echo.Context) error {
		tests, err := testHandler.testUsecase.FindAll()
		if err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
		}

		var res []responseTest
		for _, v := range tests {
			test := responseTest{
				Id: v.Id,
				Name: v.Name,
			}
			res = append(res, test)
		}

		return context.JSON(http.StatusOK, res)
	}
}

func (testHandler *testHandler) Get() echo.HandlerFunc {
	return func(context echo.Context) error {
		id, err := strconv.Atoi((context.Param("id")))
		if err != nil {
			return context.JSON(http.StatusBadRequest, err.Error())
		}

		test, err := testHandler.testUsecase.FindByID(id)
		if err != nil {
			return context.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTest{
			Id: test.Id,
			Name: test.Name,
		}

		return context.JSON(http.StatusOK, res)
	}
}

func (testHandler *testHandler) Post() echo.HandlerFunc {
	return func(context echo.Context) error {
		var req requestTest

		if err := context.Bind(&req); err != nil {
			return context.JSON(http.StatusBadRequest, err.Error())
		}

		createdTest, err := testHandler.testUsecase.Create(req.Name)
		if err != nil {
			return context.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTest{
			Id: createdTest.Id,
			Name: createdTest.Name,
		}

		return context.JSON(http.StatusCreated, res)
	}
}

func (testHandler *testHandler) Put() echo.HandlerFunc {
	return func(context echo.Context) error {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			return context.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestTest
		if err := context.Bind(&req); err != nil {
			return context.JSON(http.StatusBadRequest, err.Error())
		}

		updatedTest, err := testHandler.testUsecase.Update(id, req.Name)
		if err != nil {
			return context.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTest{
			Id:   updatedTest.Id,
			Name: updatedTest.Name,
		}

		return context.JSON(http.StatusOK, res)
	}
}

func (testHandler *testHandler) Delete() echo.HandlerFunc {
	return func(context echo.Context) error {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			return context.JSON(http.StatusBadRequest, err.Error())
		}

		err = testHandler.testUsecase.Delete(id)
		if err != nil {
			return context.JSON(http.StatusBadRequest, err.Error())
		}

		return context.NoContent(http.StatusNoContent)
	}
}