package users

import (
	"strings"

	"github.com/dung997bn/bookstore_user_api/utils/errors"
)

//User type
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

//Users type has many users
type Users []User

//Validate validates if user is valid or not
// way 1
// func Validate(user *User) *errors.RestErr {
// 	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
// 	if user.Email == "" {
// 		return errors.NewBadRequestError("Invalid email address")
// 	}
// 	return nil
// }

//Validate validates if user is valid or not
func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("Invalid password")
	}
	return nil
}
