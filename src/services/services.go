package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"models"
	"os"
	"resources"
	_ "time"
)



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

	UploadUsingConcurrency(logfiles)


	defer jsonFile.Close()
}

// UPLOAD USING CONCURRENCY

// OLD FUNCTION WITHOUT WAIT GROUPS

//var wgg sync.WaitGroup

func UploadUsingConcurrency(logfiles []models.Logfile){
	jobs := make(chan models.Logfile, len(logfiles))

	for _ ,v := range logfiles{
		jobs <- v
	}

	numberOfWorkers := 10

	for w := 1; w <= numberOfWorkers; w++ {
		//wgg.Add(1)
		go uploadHelper(jobs)
	}
	//wgg.Wait()
	close(jobs)
}


// UPLOADHELPER FUNCTION (USES GO ROUTINES TO INSERT INTO DB)

func uploadHelper(jobs <-chan models.Logfile){
	for {
		select {
			case v,ok := <-jobs:
			if(!ok) {
				return
			} else {
				//fmt.Printf("insert")
				error := resources.Db.Create(&v).Error
					if  error != nil {
						panic(error)
					}
			}
		}
	}
	//wgg.Done()
}


// FUNCTION TO GET USER RECORDS AS USER (OBFUSCATED RECORDS)

func ViewAsUser(id string, key string) []models.Logfile{

	var logfile []models.Logfile
	resources.Db.Where("" + key + " = ?", id).Find(&logfile)

	for i := 0; i < len(logfile); i++ {
		logfile[i].Email = "xxxxx@xxxx.com"
		logfile[i].Mobile = "xxxxxxxxxx"
	}

	return logfile
}

// FUNCTION TO GET USER RECORDS AS ADMIN

func ViewAsAdmin(id string, key string) []models.Logfile{

	var logfile []models.Logfile
	resources.Db.Where("" + key + " = ?", id).Find(&logfile)
	return logfile
}

