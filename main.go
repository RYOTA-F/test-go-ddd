package main

import (
	"fmt"
	gorm "tutorial-go-ddd/infrastructure/database"
)

type Test struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func main() {
    db := gorm.NewDB()
		fmt.Println("DB: 接続", db)

		var tests []Test
		db.Find(&tests)
		defer db.Close()

		fmt.Println("tests: ", tests)
}