package main

import (
	//"net/http"
	//"strconv"

	_ "database/sql"
	_"encoding/json"
	_"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_"io/ioutil"
	_"os"
	_ "time"

)

var db *gorm.DB

func init() {

	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:password@/logger?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&Logfile{}, &Logfile2{}, &Logfile3{})
}

type Logfile struct {
	gorm.Model
	Name         string   `json:"name"`
	Age          int `json:"age"`
	Email        string `json:"email"`
	Mobile		 string `json:"mobile"`
	Role         string `json:"role"`
	Num          int  `json:"num"`
	Address      string `json:"address"`
}


func main() {

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		//api.GET("/", apiHome)
		api.GET("/upload", upload)
		api.GET("/user/:key/:id", viewAsUser)
		api.GET("/admin/:key/:id", viewAsAdmin)
	}

	router.Run(":3000" )

}

func upload(c *gin.Context){

	var logfiles []Logfile
	c.BindJSON(&logfiles)

	for i := 0; i< len(logfiles) ; i++ {
		error := db.Create(&logfiles[i]).Error
		if  error != nil {
			panic(error)
		}
	}

}

func viewAsUser(c *gin.Context){
	id := c.Params.ByName("id")
	key := c.Params.ByName("key")
	var logfile []Logfile

	db.Where("" + key + " = ?", id).Find(&logfile)

	for i := 0; i < len(logfile); i++ {
		logfile[i].Email = "xxxxx@xxxx.com"
		logfile[i].Mobile = "xxxxxxxxxx"
	}

	c.JSON(200, logfile)
}


func viewAsAdmin(c *gin.Context){
	id := c.Params.ByName("id")
	key := c.Params.ByName("key")
	var logfile []Logfile

	db.Where("" + key + " = ?", id).Find(&logfile)

	c.JSON(200, logfile)
}

