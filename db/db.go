package db

import (
	"go-campaign-no-02/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	dsn := "root:password@tcp(localhost:3306)/gocompaign?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{})
}

func Manager() *gorm.DB {
	return db
}
