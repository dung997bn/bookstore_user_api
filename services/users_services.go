package services

import (
	"github.com/dung997bn/bookstore_user_api/domain/users"
	"github.com/dung997bn/bookstore_user_api/utils/errors"
)

var (
	//UsersService type
	UsersService usersServiceInterface = &usersService{}
)

//UserService type
type usersService struct{}

//UserServiceInterface type
type usersServiceInterface interface {
	GetUser(userID int64) (*users.User, *errors.RestErr)
	CreateUser(user users.User) (*users.User, *errors.RestErr)
	UpdateUser(isPatch bool, user users.User) (*users.User, *errors.RestErr)
	DeleteUser(userID int64) (int64, *errors.RestErr)
	SearchUserByStatus(status string) (users.Users, *errors.RestErr)
}

//GetUser get single User
func (u *usersService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

//CreateUser creates user
func (u *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

//UpdateUser updates existed user by Id
func (u *usersService) UpdateUser(isPatch bool, user users.User) (*users.User, *errors.RestErr) {
	currentUser, err := UsersService.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if !isPatch {
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
	} else {
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}
		if user.Email != "" {
			currentUser.Email = user.Email
		}
	}

	if err := currentUser.Validate(); err != nil {
		return nil, err
	}

	if err := currentUser.Update(); err != nil {
		return nil, err
	}
	return currentUser, nil
}

//DeleteUser deletes an exits user
func (u *usersService) DeleteUser(userID int64) (int64, *errors.RestErr) {
	currentUser, err := UsersService.GetUser(userID)
	if err != nil {
		return 0, err
	}
	if currentUser == nil {
		return 0, nil
	}

	rowsAffected, err := currentUser.Delete()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

//SearchUserByStatus finds users by status
func (u *usersService) SearchUserByStatus(status string) (users.Users, *errors.RestErr) {
	dao := users.User{Status: status}
	return dao.SearchByStatus()
}
