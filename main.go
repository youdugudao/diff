package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"contract/file"
	"contract/middleware"
	"contract/config"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin")
	})

	diff := router.Group("/diff")
	diff.Use(middleware.AuthRequire)
	diff.POST("", file.PostDiff)
	diff.POST("/one", file.PostDiffOne)

	router.Run(":" + config.Port)
}
