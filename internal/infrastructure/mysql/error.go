package mysql

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

const (
	ErrDuplicateEntry = 1062
)

func IsDuplicateError(err error) bool {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return mysqlErr.Number == ErrDuplicateEntry
	}
	return false
}
