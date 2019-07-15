package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"models"
	"resources"
)

// UPLOAD FUNCTION (PARSE JSON FILES AND INSERT IN DB)

func Upload(c *gin.Context){

	var logfiles []models.Logfile
	c.BindJSON(&logfiles)

	jobs := make(chan int, 100)

	for w := 1; w <= 5; w++ {
		go uploadHelper(w, logfiles, jobs)
	}
	for j := 0; j < len(logfiles); j++ {
		jobs <- j
	}
	close(jobs)
}

// UPLOADHELPER FUNCTION (USES GO ROUTINES TO INSERT INTO DB)

func uploadHelper(w int,logfiles []models.Logfile, jobs <-chan int){
	for j := range jobs {
		fmt.Println("worker ",w," started  insert", j)
		error := resources.Db.Create(&logfiles[j]).Error
		if  error != nil {
			panic(error)
		}
		fmt.Println("worker ",w," finished insert", j)
	}
}

// FUNCTION TO GET USER RECORDS AS USER (OBFUSCATED RECORDS)


func ViewAsUser(c *gin.Context){
	id := c.Params.ByName("id")
	key := c.Params.ByName("key")
	var logfile []models.Logfile

	resources.Db.Where("" + key + " = ?", id).Find(&logfile)

	for i := 0; i < len(logfile); i++ {
		logfile[i].Email = "xxxxx@xxxx.com"
		logfile[i].Mobile = "xxxxxxxxxx"
	}

	c.JSON(200, logfile)
}

// FUNCTION TO GET USER RECORDS AS ADMIN

func ViewAsAdmin(c *gin.Context){
	id := c.Params.ByName("id")
	key := c.Params.ByName("key")
	var logfile []models.Logfile

	resources.Db.Where("" + key + " = ?", id).Find(&logfile)

	c.JSON(200, logfile)
}
