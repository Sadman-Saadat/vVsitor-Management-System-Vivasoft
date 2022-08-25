package database

import (
	"fmt"
<<<<<<< HEAD
	//"github.com/jinzhu/gorm"
	"gorm.io/driver/sqlite"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
	//	"visitor-management-system/config"
=======
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// "visitor-management-system/config"
	//"database/sql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	//gormlogger "gorm.io/gorm/logger"
>>>>>>> dad55e8260aedf1d4ad7f78775d3ad4da2c70dee
	"visitor-management-system/model"
)

var db *gorm.DB

// func Connect() {
// 	dns := fmt.Sprintf("%s/%s?charset=utf8mb4&parseTime=True", config.GetConfig().SqlUri, config.GetConfig().SqlDb)
// 	database, err := gorm.Open("mysql", dns)

// 	if err != nil {
// 		fmt.Println("error connecting to db")
// 		panic(err)
// 	} else {
// 		fmt.Println("connected to db")
// 	}
// 	db = database
// }

func Connect() {
<<<<<<< HEAD
	database, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
=======
	//logMode := gormlogger.Info

	database, err := gorm.Open(sqlite.Open("/app/gorm/gorm.db"), &gorm.Config{
		//PrepareStmt: true,
		//Logger:      gormlogger.Default.LogMode(logMode),
	})
>>>>>>> dad55e8260aedf1d4ad7f78775d3ad4da2c70dee

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
	db.AutoMigrate(&model.Branch{})
	db.AutoMigrate(&model.User{})
	//db.AutoMigrate(&model.Subscription{})
	db.AutoMigrate(&model.MasterAdmin{})
	db.AutoMigrate(&model.Package{})
	db.AutoMigrate(&model.PackageFeatures{})
	db.AutoMigrate(&model.Visitor{})
	db.AutoMigrate(&model.TrackVisitor{})
	db.AutoMigrate(&model.Record{})
<<<<<<< HEAD

	fmt.Println("all migrated")
	//db.Model(&model.TrackVisitor{}).AddForeignKey("v_id", "visitors(id)", "RESTRICT", "RESTRICT")

	//db.Model(&model.Subscription{}).AddForeignKey("company_id", "companies(id)", "RESTRICT", "RESTRICT")
=======
	db.AutoMigrate(&model.Setting{})
	db.AutoMigrate(&model.UserBranchRelation{})
	// db.Model(&model.TrackVisitor{}).AddForeignKey("v_id", "visitors(id)", "RESTRICT", "RESTRICT")
	// db.Model(&model.Subscription{}).AddForeignKey("company_id", "companies(id)", "RESTRICT", "RESTRICT")
	// db.Model(&model.User{}).AddForeignKey("branch_id", "branchs(id)", "RESTRICT", "RESTRICT")
>>>>>>> dad55e8260aedf1d4ad7f78775d3ad4da2c70dee

}
func GetDB() *gorm.DB {
	if db == nil {
		Connect()
	}

	return db
}
