package user

import (
	"net/http"
	"strconv"

	"github.com/dung997bn/bookstore_user_api/domain/users"
	"github.com/dung997bn/bookstore_user_api/services"
	"github.com/dung997bn/bookstore_user_api/utils/errors"
	"github.com/gin-gonic/gin"
)

//GetUser gets single user
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("User id shoud be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
	}
	c.JSON(http.StatusOK, user)
}

//CreateUser creates an user
func CreateUser(c *gin.Context) {
	var user users.User
	// way 1 to read data from request body
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	//Handle error
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	//Handle json error
	// 	return
	// }

	//way 2
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		//Handle error
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//Handle save error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

//SearchUser search users by condition
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me")
}
