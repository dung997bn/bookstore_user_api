package users

import (
	"fmt"

	"github.com/dung997bn/bookstore_user_api/logger"

	"github.com/dung997bn/bookstore_user_api/datasources/mysql/usersdb"
	"github.com/dung997bn/bookstore_user_api/utils/cryptoutils"
	"github.com/dung997bn/bookstore_user_api/utils/dateutils"
	"github.com/dung997bn/bookstore_user_api/utils/errors"
	"github.com/dung997bn/bookstore_user_api/utils/mysqlutils"
)

const (
	//StatusActive type
	StatusActive = "active"
	//StatusBlock type
	StatusBlock = "block"
)

const (
	//query
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?,?,?,?,?,?) ;"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status From users Where id=? ; "
	queryUpdateUser             = "Update users set first_name=?, last_name =?, email =?, date_created=?, status=?, password=? Where id=?;"
	queryDeleteUser             = "Delete from users Where id=?;"
	queryFindUserByStatus       = "Select id, first_name, last_name, email, date_created, status From users Where status=? ;"
	queryFindByEmailAndPassword = "Select id, first_name, last_name, email, date_created, status from users Where email=? and password=? and status=? ;"
)

//Get gets single by id
func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error when trying to get user by id", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

//Save user
func (user *User) Save() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = dateutils.GetNowDBFormat()
	user.Status = StatusActive
	user.Password = cryptoutils.GetMD5(user.Password)
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return mysqlutils.ParseError(saveErr)
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to save user", err)
		return mysqlutils.ParseError(err)
	}
	user.ID = userID
	return nil
}

//Update func
func (user *User) Update() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = dateutils.GetNowDBFormat()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password, user.ID)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return mysqlutils.ParseError(err)
	}
	return nil
}

//Delete func
func (user *User) Delete() (int64, *errors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return 0, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Exec(user.ID)
	if err != nil {
		logger.Error("error when trying to delete user", err)
		return 0, mysqlutils.ParseError(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("error when trying to delete user", err)
		return 0, mysqlutils.ParseError(err)
	}
	return rowsAffected, nil
}

//SearchByStatus func
func (user *User) SearchByStatus() ([]User, *errors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare search user statement", err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(user.Status)
	if err != nil {
		logger.Error("error when trying to search user", err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when trying to search user", err)
			return nil, mysqlutils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		logger.Error("no users matching", nil)
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching with status: %s", user.Status))
	}
	return results, nil
}

//FindByEmailAndPassword gets single user by email & pass
func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare find user by email & password statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	fmt.Println(user.Email)
	result := stmt.QueryRow(user.Email, cryptoutils.GetMD5(user.Password), StatusActive)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error when trying to get user by email & password", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}
