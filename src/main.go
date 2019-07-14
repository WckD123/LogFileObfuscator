package main

import (
	//"net/http"
	//"strconv"

	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/static"
	_ "github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"os"
	_ "time"

	//"net/http"
)

var db *gorm.DB

func init() {

	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:password@/logger?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	// Close the DB connection after everything is done.
	defer db.Close()

	//Migrate the schema
	db.AutoMigrate(&Logfile{}, &Logfile2{}, &Logfile3{})
}

type Logfile struct {
	gorm.Model
	Name         string   `json:"name"`
	Age          int `json:"age"`
	Email        string `json:"email"`
	Mobile		 int `json:"mobile"`
	Role         string `json:"role"`
	Num          int  `json:"num"`
	Address      string `json:"address"`
}

type Logfile2 struct {
	gorm.Model
	Name         string `json:"name"`
	Age          int `json:"age"`
	Email        string `json:"email"`
	Mobile		 int `json:"mobile"`
	School         string `json:"school"`
	Num          int `json:"num"`
	College      string `json:"college"`
}

type Logfile3 struct {
	gorm.Model
	Name         string `json:"name"`
	Age          int `json:"age"`
	Email        string `json:"email"`
	Mobile		 int `json:"mobile"`
	Company         string  `json:"company"`
	Num          int `json:"num"`
}


func main() {

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		//api.GET("/", apiHome)
		api.GET("/filePath", upload)
		api.GET("/user/:key/:id", viewAsUser)
		api.GET("/admin/:id", viewAsAdmin)
	}

	router.Run(":3001" )

}

func upload(c *gin.Context){

}

func viewAsUser(c *gin.Context){

}


func viewAsAdmin(c *gin.Context){
	id := c.Params.ByName("id")
	key := c.Params.ByName("key")
	var logfile []Logfile

	db.Where("" + key + " = ?", id).Find(&logfile)

	c.JSON(200, logfile)
}


func readJSON(path string){
	var logfile []Logfile

	json, err := os.Open(path)
	defer json.Close()

	if err != nil {
		panic(err.Error())
	}

	byteValue, _ := ioutil.ReadAll(json)
	json.Unmarshal(byteValue, &logfile)

}

