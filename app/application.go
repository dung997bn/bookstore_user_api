package app

import (
	"github.com/dung997bn/bookstore_user_api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplication starts application
func StartApplication() {
	mapUrls()

	logger.Info("about to start the application...")
	router.Run(":8081")

}
