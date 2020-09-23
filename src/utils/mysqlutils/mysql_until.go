package mysqlutils

import (
	"errors"
	"strings"

	"github.com/dung997bn/bookstore_utils-go/resterrors"
	"github.com/go-sql-driver/mysql"
)

const (
	noRowsError = "no rows in result set"
)

//ParseError parse error to Mysql error
func ParseError(err error) resterrors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), noRowsError) {
			return resterrors.NewNotFoundError("No record matching with given parameter")
		}
		return resterrors.NewInternalServerError("Error when process request", errors.New("error mysql server"))
	}

	switch sqlErr.Number {
	case 1062:
		return resterrors.NewBadRequestError("Invalid data")
	}
	return resterrors.NewInternalServerError("Error when process request", errors.New("error mysql server"))
}
