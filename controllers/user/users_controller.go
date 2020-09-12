package user

import (
	"net/http"
	"strconv"

	"github.com/dung997bn/bookstore_user_api/domain/users"
	"github.com/dung997bn/bookstore_user_api/services"
	"github.com/dung997bn/bookstore_user_api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getHeaderPublic(header string) bool {
	isPublic, err := strconv.ParseBool(header)
	if err != nil {
		return false
	}
	return isPublic
}

func getUserIDFromParam(userIDParam string) (int64, *errors.RestErr) {
	userID, IDErr := strconv.ParseInt(userIDParam, 10, 64)
	if IDErr != nil {
		return 0, errors.NewBadRequestError("User id shoud be a number")
	}
	return userID, nil
}

//Get gets single user
func Get(c *gin.Context) {
	userID, IDErr := getUserIDFromParam(c.Param("user_id"))
	if IDErr != nil {
		c.JSON(IDErr.Status, IDErr)
		return
	}

	user, getErr := services.UsersService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(getHeaderPublic(c.GetHeader("X-public"))))
}

//Create creates an user
func Create(c *gin.Context) {
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
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(getHeaderPublic(c.GetHeader("X-public"))))
}

//Update updates existed user by Id
func Update(c *gin.Context) {
	userID, IDErr := getUserIDFromParam(c.Param("user_id"))
	if IDErr != nil {
		c.JSON(IDErr.Status, IDErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.ID = userID

	isPatch := c.Request.Method == http.MethodPatch

	result, UpdateErr := services.UsersService.UpdateUser(isPatch, user)
	if UpdateErr != nil {
		c.JSON(UpdateErr.Status, UpdateErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(getHeaderPublic(c.GetHeader("X-public"))))
}

//Delete deletes an exits user
func Delete(c *gin.Context) {
	userID, IDErr := getUserIDFromParam(c.Param("user_id"))
	if IDErr != nil {
		c.JSON(IDErr.Status, IDErr)
		return
	}
	rowsAffected, err := services.UsersService.DeleteUser(userID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, rowsAffected)
}

//SearchByStatus finds users by status
func SearchByStatus(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.SearchUserByStatus(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, users.Marshall(getHeaderPublic(c.GetHeader("X-public"))))
}
