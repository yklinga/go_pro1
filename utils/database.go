package utils

import (
	"go_pro1/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func BaseDB() *gorm.DB  {
	dsn := "root:Root.123@tcp(103.46.128.49:52706)/gin_database?charset=utf8&parseTime=True"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database, err" + err.Error())
	}
	db.AutoMigrate(&model.User{})

	DB = db
	return db
}

func GetDB() * gorm.DB  {
	return DB
}