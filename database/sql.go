package database

import (
	"fmt"
	//"github.com/jinzhu/gorm"
	"gorm.io/driver/sqlite"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
	//	"visitor-management-system/config"
	"visitor-management-system/model"
)

var db *gorm.DB

func Connect() {
	database, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("error connecting to db")
		panic(err)
	} else {
		fmt.Println("connected to db")
	}
	db = database
}

func Migration() {
	db.Exec("PRAGMA foreign_keys = ON")
	db.AutoMigrate(&model.Company{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Subscription{})
	db.AutoMigrate(&model.Visitor{})
	db.AutoMigrate(&model.TrackVisitor{})
	db.AutoMigrate(&model.Record{})
	//db.Model(&model.TrackVisitor{}).AddForeignKey("v_id", "visitors(id)", "RESTRICT", "RESTRICT")

	//db.Model(&model.Subscription{}).AddForeignKey("company_id", "companies(id)", "RESTRICT", "RESTRICT")

}
func GetDB() *gorm.DB {
	if db == nil {
		Connect()
	}

	return db
}
