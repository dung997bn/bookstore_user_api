package services

import (
	"github.com/dung997bn/bookstore_user_api/domain/users"
	"github.com/dung997bn/bookstore_user_api/utils/errors"
)

//GetUser get single User
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	if userID <= 0 {
		return nil, errors.NewBadRequestError("Invalid user id")
	}
	return nil, nil
}

//CreateUser creates user
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
