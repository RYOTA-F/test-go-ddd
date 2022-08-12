package main

import (
	infra "tutorial-go-ddd/infrastructure"
	database "tutorial-go-ddd/infrastructure/database"
	"tutorial-go-ddd/interface/handler"
	"tutorial-go-ddd/usecase"

	"github.com/labstack/echo"
)

func main() {
	testRepository := infra.NewTestRepotitory(database.NewDB())
	testUsecase := usecase.NewTestUseCase(testRepository)
	testHandler := handler.NewTestHandler(testUsecase)

	e := echo.New()
	handler.InitRouting(e, testHandler)
	e.Logger.Fatal(e.Start(":3000"))
}