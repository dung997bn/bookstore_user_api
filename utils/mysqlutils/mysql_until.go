package mysqlutils

import (
	"strings"

	"github.com/dung997bn/bookstore_user_api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	noRowsError = "no rows in result set"
)

//ParseError parse error to Mysql error
func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), noRowsError) {
			return errors.NewNotFoundError("No record matching with given parameter")
		}
		return errors.NewInternalServerError("Error when process request")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("Invalid data")
	}
	return errors.NewInternalServerError("Error when process request")
}
