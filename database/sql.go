package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"visitor-management-system/config"
	"visitor-management-system/model"
)

var db *gorm.DB

func Connect() {
	dns := fmt.Sprintf("%s/%s?charset=utf8mb4&parseTime=True", config.GetConfig().SqlUri, config.GetConfig().SqlDb)
	database, err := gorm.Open("mysql", dns)

	if err != nil {
		fmt.Println("error connecting to db")
		panic(err)
	} else {
		fmt.Println("connected to db")
	}
	db = database
}

func Migration() {
	db.AutoMigrate(&model.Company{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Subscription{})
	db.AutoMigrate(&model.Visitor{})
	db.AutoMigrate(&model.TrackVisitor{})
	db.AutoMigrate(&model.Record{})
	db.Model(&model.TrackVisitor{}).AddForeignKey("v_id", "visitors(id)", "RESTRICT", "RESTRICT")

	db.Model(&model.Subscription{}).AddForeignKey("company_id", "companies(id)", "RESTRICT", "RESTRICT")

}
func GetDB() *gorm.DB {
	if db == nil {
		Connect()
	}

	return db
}
