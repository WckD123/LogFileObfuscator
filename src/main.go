package main

import (
	_ "database/sql"
	_ "encoding/json"
	"fmt"
	_ "fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "io/ioutil"
	"time"
	_ "time"
)

var db *gorm.DB

func init() {

	// CREATE DB CONNECTION
	var err error
	db, err = gorm.Open("mysql", "root:password@/logger?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	// AUTOMIGRATE THE SCHEMA
	db.AutoMigrate(&Logfile{})
}

// DECLARING MODEL

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

// UPLOAD FUNCTION (PARSE JSON FILES AND INSERT IN DB)

func upload(c *gin.Context){

	var logfiles []Logfile
	c.BindJSON(&logfiles)

	jobs := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go uploadHelper(logfiles, jobs)
	}

	for j := 0; j < len(logfiles); j++ {
		jobs <- j
	}
	close(jobs)
}

func uploadHelper(logfiles []Logfile, jobs <-chan int){
	for j := range jobs {
		fmt.Println("worker started  job", j)
		error := db.Create(&logfiles[j]).Error
		if  error != nil {
			panic(error)
		}
		fmt.Println("worker finished job", j)
	}
}

// FUNCTION TO GET USER RECORDS AS USER (OBFUSCATED RECORDS)

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

// FUNCTION TO GET USER RECORDS AS ADMIN

func viewAsAdmin(c *gin.Context){
	id := c.Params.ByName("id")
	key := c.Params.ByName("key")
	var logfile []Logfile

	db.Where("" + key + " = ?", id).Find(&logfile)

	c.JSON(200, logfile)
}

