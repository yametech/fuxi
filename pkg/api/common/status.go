package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	PageKey  = "page"
	PageSize = "pageSize"
)

const (
	msg    = "message"
	data   = "data"
	status = "errors"
	code   = "code"
)

func ToRequestParamsError(g *gin.Context, err error) {
	g.JSON(
		http.StatusBadRequest,
		gin.H{
			code:   http.StatusBadRequest,
			data:   "",
			msg:    fmt.Sprintf("%s: %s", "Status bad request", err.Error()),
			status: err.Error(),
		},
	)
}

func ResourceNotFoundError(g *gin.Context, runtimeData interface{}, err error) {
	g.JSON(
		http.StatusOK,
		gin.H{
			code:   http.StatusOK,
			data:   runtimeData,
			msg:    fmt.Sprintf("%s: %s", "Status internal server error", err.Error()),
			status: err.Error(),
		},
	)
}

func ToInternalServerError(g *gin.Context, runtimeData interface{}, err error) {
	g.JSON(
		http.StatusInternalServerError,
		gin.H{
			code:   http.StatusInternalServerError,
			data:   runtimeData,
			msg:    fmt.Sprintf("%s: %s", "Status internal server error", err.Error()),
			status: err.Error(),
		},
	)
}
