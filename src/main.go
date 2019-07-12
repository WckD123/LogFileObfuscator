package main

import (
	//"net/http"
	//"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:password@/logger?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	//
	////Migrate the schema
	//db.AutoMigrate(&todoModel{})
}

func main() {
	router := gin.Default()

	router.Run()

}
