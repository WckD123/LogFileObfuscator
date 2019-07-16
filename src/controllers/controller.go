package controllers

import (
	"github.com/gin-gonic/gin"
	"models"
	"services"
)

func Upload(c *gin.Context) {
	var logfiles []models.Logfile
	c.BindJSON(&logfiles)
	services.UploadUsingConcurrency(logfiles)
}

func ViewAsUser(c *gin.Context){
	id := c.Params.ByName("id")
	key := c.Params.ByName("key")
	logfile := services.ViewAsUser(id, key);

	c.JSON(200, logfile)
}

func ViewAsAdmin(c *gin.Context){
	id := c.Params.ByName("id")
	key := c.Params.ByName("key")
	logfile := services.ViewAsAdmin(id, key);

	c.JSON(200, logfile)
}

func UploadWithPath(c *gin.Context){
	path := c.Params.ByName("path")
	services.UploadWithPath(path)
}