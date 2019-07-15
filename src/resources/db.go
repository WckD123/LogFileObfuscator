package resources

import (
	"github.com/jinzhu/gorm"
	"models"
)

var Db *gorm.DB

// CREATE DB CONNECTION

func ConnectDB(){
	var err error
	Db, err = gorm.Open("mysql", "root:password@/logger?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
}

// AUTOMIGRATE THE SCHEMA

func MigrateDB(){
	Db.AutoMigrate(&models.Logfile{})
}
