package controllers

import (
	"net/http"
	"quickstart/database"

	"github.com/gin-gonic/gin"
)




func GetBookStats(c *gin.Context){
	booklen:=len(database.Books)
	
	c.JSON(http.StatusOK,gin.H{
		"total number":booklen,
	})
}



