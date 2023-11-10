package response

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SignupSuccess(c *gin.Context, code int, userid uuid.UUID) {
	c.JSON(code, gin.H{
		"status":  "success",
		"userid":  userid,
	})
}

func Success(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status":  "error",
		"message": message,
	})
}