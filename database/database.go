package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	const MYSQL = "root@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	dns := MYSQL
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Cannot Connect to Database!")
	}

	fmt.Println("Connected to Database!")
}
