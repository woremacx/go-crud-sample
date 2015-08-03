package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/woremacx/go-crud-sample/models"
)

func main() {
	db, _ := gorm.Open("mysql", "root:@/testdb?charset=utf8&parseTime=True")
	db.CreateTable(&models.User{})
}
