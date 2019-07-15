package controllers

import (
	"github.com/gin-gonic/gin"
	_ "models"
	"services"
)

func Upload(c *gin.Context) {
	services.Upload(c)
}

func ViewAsUser(c *gin.Context){
	services.ViewAsUser(c)
}

func ViewAsAdmin(c *gin.Context){
	services.ViewAsAdmin(c);
}

func UploadWithPath(c *gin.Context){
	path := c.Params.ByName("path")
	services.UploadWithPath(path)
}