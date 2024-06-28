package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"fmt"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/b_apigo"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	database.AutoMigrate(&User{})

	DB = database
}