package app

import (
	"github.com/dung997bn/bookstore_user_api/controllers/ping"
	"github.com/dung997bn/bookstore_user_api/controllers/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", user.GetUser)
	router.POST("/users", user.CreateUser)
}
