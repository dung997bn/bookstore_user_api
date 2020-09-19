package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dung997bn/bookstore_oauth-go/oauth"
	"github.com/dung997bn/bookstore_user_api/domain/users"
	"github.com/dung997bn/bookstore_user_api/services"
	"github.com/dung997bn/bookstore_utils-go/resterrors"
	"github.com/gin-gonic/gin"
)

func getHeaderPublic(header string) bool {
	isPublic, err := strconv.ParseBool(header)
	if err != nil {
		return false
	}
	return isPublic
}

func getUserIDFromParam(userIDParam string) (int64, *resterrors.RestErr) {
	userID, IDErr := strconv.ParseInt(userIDParam, 10, 64)
	if IDErr != nil {
		return 0, resterrors.NewBadRequestError("User id shoud be a number")
	}
	return userID, nil
}

//Get gets single user
func Get(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

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
	if oauth.GetCallerID(c.Request) != user.ID {

		c.JSON(http.StatusOK, user.Marshall(false))
		return
	}
	c.JSON(http.StatusOK, user.Marshall(oauth.IsPublic(c.Request)))
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
		restErr := resterrors.NewBadRequestError("Invalid json body")
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
		restErr := resterrors.NewBadRequestError("Invalid json body")
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

//Login func
func Login(c *gin.Context) {
	var request users.LoginRequest
	fmt.Println(request)
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, err := services.UsersService.LoginUser(&request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(getHeaderPublic(c.GetHeader("X-public"))))
}
