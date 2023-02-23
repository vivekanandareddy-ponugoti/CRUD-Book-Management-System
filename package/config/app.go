package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.db
)

func Connect() {
	d, err := gorm.Open("mysql", "vivekreddy13:1q2w3e4r/book-store?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.db {
	return db
}
