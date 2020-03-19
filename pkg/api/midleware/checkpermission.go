package midleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/db"
)

type HasPermissionFunc func(uint32) bool

func CheckPermission(hasPermissionFunc func(uint32) bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		userName := c.Request.Header.Get("x-auth-username")
		//check current user
		user, err := db.FindUserByName(userName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err})
			return
		}
		val, err := user.FindCurrentUserPermissionValue()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err})
			return
		}
		if hasPermissionFunc(uint32(val)) {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "Unauthorized"})
			return
		}
	}

}
