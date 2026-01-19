package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func AuthRequired() gin.HandlerFunc{
	return func(c *gin.Context){
		apiKey := c.GetHeader("X-API-Key")
 if apiKey !="secret123"{
	c.JSON(http.StatusUnauthorized,gin.H{
		"error":"Unauthorized - INvalid API key",
	})
	c.Abort()
	return
 }
 c.Next()

	}
}