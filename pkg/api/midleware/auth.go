package midleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Auth
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userName := c.Request.Header.Get("x-auth-username")
		if len(userName) < 1 {
			c.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"err": "authorization not correct"})
			return
		} else {
			c.Next()
		}
	}

}
