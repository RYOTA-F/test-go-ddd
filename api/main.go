package main

import (
	infra "tutorial-go-ddd/infrastructure"
	handler "tutorial-go-ddd/interface"
	"tutorial-go-ddd/usecase"

	"github.com/labstack/echo"
)

func main() {
	testRepository := infra.NewTestRepotitory(infra.NewDB())
	testUsecase := usecase.NewTestUseCase(testRepository)
	testHandler := handler.NewTestHandler(testUsecase)

	e := echo.New()
	handler.Router(e, testHandler)
	e.Logger.Fatal(e.Start(":3000"))
}