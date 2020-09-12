package app

import (
	"github.com/dung997bn/bookstore_user_api/controllers/ping"
	"github.com/dung997bn/bookstore_user_api/controllers/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", user.Get)
	router.GET("/internal/users/search", user.SearchByStatus)
	router.POST("/users", user.Create)
	router.PUT("/users/:user_id", user.Update)
	router.PATCH("/users/:user_id", user.Update)
	router.DELETE("/users/:user_id", user.Delete)

}
