package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)


func Logger() gin.HandlerFunc {
	return func(c *gin.Context){
		starttime := time.Now()

		c.Next()

		duration :=time.Since(starttime)
		statusCode :=c.Writer.Status()

		fmt.Printf("[%s] %s %s | Status: %d | Duration: %v\n",
            time.Now().Format("2006-01-02 15:04:05"),
            c.Request.Method,
            c.Request.URL.Path,
            statusCode,
            duration,
        )

	}
}