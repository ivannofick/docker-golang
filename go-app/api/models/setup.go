package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/api_golang"))
	if err != nil {
		fmt.Println("Gagal Terkoneksi Database")
	}
	db.AutoMigrate(&User{})
	DB = db
}
