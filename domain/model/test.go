package model

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

type Test struct {
  Id uint
  name string
  CreatedAt time.Time
}