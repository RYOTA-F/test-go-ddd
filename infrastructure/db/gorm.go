package db

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	gorm "github.com/jinzhu/gorm"
	godotnev "github.com/joho/godotenv"
)

// NewDB DBと接続する
func NewDB() *gorm.DB {
	// ENV読み取り
	err := godotnev.Load(".env")
	if err != nil {
		fmt.Printf("読み込みに失敗しました: %v", err)
	}

	// 接続情報を設定
	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	// (localhost:3306ではなく) (コンテナ名:3306)
	HOST := "tcp(db:3306)"
	DBNAME := os.Getenv("DB_NAME")

	CONNECT := USER + ":" + PASS + "@" + HOST + "/" + DBNAME

	// DBに接続
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	return db
}
