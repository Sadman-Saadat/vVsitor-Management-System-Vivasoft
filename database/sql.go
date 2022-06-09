package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect(database_name string) {
	dns := fmt.Sprintf("root:quddus1916@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", database_name)
	database, err := gorm.Open("mysql", dns)

	if err != nil {
		fmt.Println("error connecting to db")
		panic(err)
	} else {
		fmt.Println("connected to db")
	}
	db = database
}

func GetDB(db_name string) *gorm.DB {
	Connect(db_name)

	return db
}
