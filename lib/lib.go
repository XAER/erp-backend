package lib

import "github.com/gin-gonic/gin"

func RespondWithError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, gin.H{
		"error": message,
	})
}
