package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "TBD",
	})
}
