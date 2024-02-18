package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"newsSharing/src/config"
	"newsSharing/src/fetchNews"
	"newsSharing/src/guardian"
	"newsSharing/src/nyTimes"
)

func init() {
	config.DoInit()
	guardian.DoInit()
	nyTimes.DoInit()
}

func main() {
	mainRouter := gin.Default()
	mainRouter.GET("/ping", func(ginContext *gin.Context) {
		ginContext.JSON(http.StatusOK, gin.H{
			"message": "pong !!!!",
		})
	})

	mainRouter.GET("/fetchNews", fetchNews.NewsController)
	mainRouter.Run("0.0.0.0:" + "8080")
}
