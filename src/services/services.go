package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"models"
	"os"
	"resources"
	"time"
	_ "time"
)

// UPLOAD FUNCTION

func Upload(c *gin.Context){

	var logfiles []models.Logfile
	c.BindJSON(&logfiles)

}


// UPLOAD WITH PATH FUNCTION

func UploadWithPath(path string){
	var logfiles []models.Logfile
	jsonFile, err := os.Open(path)
	if err != nil{
		panic(err)
	}

	fmt.Println("JSON file opened")

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &logfiles)

	t1 := time.Now()

	UploadUsingConcurrency(logfiles)

	t2 := time.Now()

	diff := t2.Sub(t1)
	fmt.Println(diff)

	defer jsonFile.Close()
}

// UPLOAD USING CONCURRENCY

// OLD FUNCTION WITHOUT WAIT GROUPS

func UploadUsingConcurrency(logfiles []models.Logfile){
	jobs := make(chan models.Logfile, len(logfiles))

	for _ ,v := range logfiles{
		jobs <- v
	}

	numberOfWorkers := 30

	for w := 1; w <= numberOfWorkers; w++ {
		go uploadHelper(jobs)
	}
	close(jobs)
}

// UPLAOD FUNCTION WITH WAIT GROUP (SLOW)

//var wg sync.WaitGroup
//
//func UploadUsingConcurrency(logfiles []models.Logfile){
//	//jobs := make(chan int, len(logfiles))
//	wg.Add(len(logfiles))
//	for _, logfile := range logfiles {
//		go func(logfile1 models.Logfile){
//			error := resources.Db.Create(&logfile).Error
//			if  error != nil {
//				panic(error)
//			}
//			wg.Done()
//		}(logfile)
//	}
//	wg.Wait()
//}

// UPLOADHELPER FUNCTION (USES GO ROUTINES TO INSERT INTO DB)

func uploadHelper(jobs <-chan models.Logfile){
	for {
		select {
			case v,ok := <-jobs:
			if(!ok) {
				return
			} else {
				error := resources.Db.Create(&v).Error
					if  error != nil {
						panic(error)
					}
			}
		}
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

