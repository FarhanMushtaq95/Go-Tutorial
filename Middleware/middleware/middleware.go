package middleware

import "github.com/gin-gonic/gin"

func Authenticate(c *gin.Context) {
	if !(c.Request.Header.Get("token") == "auth") {
		c.AbortWithStatusJSON(500, gin.H{
			"Message": "Token not found",
		})
		return
	}
	c.Next()
}
