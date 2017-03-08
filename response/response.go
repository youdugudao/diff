package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type resData struct {
	Success     bool
	SuccessInfo interface{}
	ErrorInfo   interface{}
}

func Ok(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, resData{true, data, nil})
}

func Error(errInfo string, c *gin.Context) {
	c.JSON(http.StatusOK, resData{false, nil, errInfo})
}
