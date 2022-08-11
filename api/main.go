package main

import (
	"fmt"

	"tutorial-go-ddd/infrastructure/db"
)

type Test struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func main() {
    db := db.NewDB()
		fmt.Println("DB: 接続", db)

		var tests []Test
		db.Find(&tests)
		defer db.Close()

		fmt.Println("tests: ", tests)
}