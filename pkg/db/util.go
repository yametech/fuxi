package db

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) uint32 {
	// jwt rewrite header
	tokenUser := c.Request.Header.Get("x-auth-username")
	if len(tokenUser) < 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"err": "authorization not correct"})
		return 0
	}

	return 1

}
